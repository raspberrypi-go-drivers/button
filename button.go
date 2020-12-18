package button

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
)

// State represents the current state of a button
type State bool

const (
	fetchDelta time.Duration = 10 * time.Millisecond
	// Pressed is the state of a pressed button
	Pressed State = true
	// Released is the state of a released button
	Released State = false
)

// Button instance
type Button struct {
	pinID         uint8
	pin           rpio.Pin
	edgeDetection bool
	EventChan     chan State
}

// NewButton creates a new Button instance
// pinID is the GPIO pin number you're using
func NewButton(pinID uint8) *Button {
	button := Button{
		pinID:         pinID,
		edgeDetection: false,
	}
	button.pin = rpio.Pin(button.pinID)
	button.pin.Mode(rpio.Input)
	button.pin.PullDown()
	return &button
}

// SetPullDown set the GPIO to PullDown (GPIO to Vcc)
func (button *Button) SetPullDown() error {
	button.pin.PullDown()
	return nil
}

// SetPullUp set the GPIO to PullUp (GPIO to GND)
func (button *Button) SetPullUp() error {
	button.pin.PullUp()
	return nil
}

// GetState return the current state of the Button
func (button *Button) GetState() (State, error) {
	state := button.pin.Read()
	if state == 1 {
		return true, nil
	} else if state == 0 {
		return false, nil
	}
	return false, fmt.Errorf("unknown state '%d'", state)
}

// EnableEventDetection enables event detection
func (button *Button) EnableEventDetection() error {
	previousState, err := button.GetState()
	if err != nil {
		return err
	}
	button.EventChan = make(chan State, 1)
	go func() {
		for {
			currentState, err := button.GetState()
			if err != nil {
				log.WithField("error", err).Error("impossible to get state")
			}
			if currentState != previousState {
				button.EventChan <- currentState
				previousState = currentState
			}
			time.Sleep(fetchDelta)
		}
	}()
	return nil
}
