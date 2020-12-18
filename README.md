# Button

[![PkgGoDev](https://pkg.go.dev/badge/github.com/bbayszczak/raspberrypi-go-drivers/led)](https://pkg.go.dev/github.com/bbayszczak/raspberrypi-go-drivers/button)

This drivers allows interact with a button connected to the GPIO

## Documentation

For full documentation, please visit [![PkgGoDev](https://pkg.go.dev/badge/github.com/bbayszczak/raspberrypi-go-drivers/led)](https://pkg.go.dev/github.com/bbayszczak/raspberrypi-go-drivers/button)

## Quick start

```go
import (
	"fmt"

	"github.com/raspberrypi-go-drivers/button"
	"github.com/stianeikeland/go-rpio/v4"
	log "github.com/sirupsen/logrus"
)

func main() {
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
```
