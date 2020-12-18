package button_test

import (
	"fmt"
	"os"
	"time"

	"github.com/raspberrypi-go-drivers/button"
	log "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
)

// Display button status each time the button state changes
func Example() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	b1 := button.NewButton(17)
	if err := b1.EnableEventDetection(); err != nil {
		log.WithField("error", err).Error("impossible to enable edge detection")
	}
	for event := range b1.EventChan {
		if event == button.Pressed {
			fmt.Println("Button pressed")
		} else if event == button.Released {
			fmt.Println("Button released")
		}
	}
}

func ExampleNewButton() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	b1 := button.NewButton(17)
	b1State, _ := b1.GetState()
	fmt.Println(b1State)
}

func ExampleButton_SetPullDown() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	b1 := button.NewButton(17)
	if err := b1.SetPullDown(); err != nil {
		log.WithField("error", err).Error("impossible to set pulldown")
	}
}

func ExampleButton_SetPullUp() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	b1 := button.NewButton(17)
	if err := b1.SetPullUp(); err != nil {
		log.WithField("error", err).Error("impossible to set pullup")
	}
}

func ExampleButton_GetState() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	b1 := button.NewButton(17)
	for {
		state, _ := b1.GetState()
		fmt.Println(state)
		time.Sleep(time.Second)
	}
}

func ExampleButton_EnableEventDetection() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	b1 := button.NewButton(17)
	if err := b1.EnableEventDetection(); err != nil {
		log.WithField("error", err).Error("impossible to enable edge detection")
	}
	for event := range b1.EventChan {
		if event == button.Pressed {
			fmt.Println("Button pressed")
		} else if event == button.Released {
			fmt.Println("Button released")
		}
	}
}
