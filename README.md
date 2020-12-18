# Button

[![Go Reference](https://pkg.go.dev/badge/github.com/raspberrypi-go-drivers/button.svg)](https://pkg.go.dev/github.com/raspberrypi-go-drivers/button)
![golangci-lint](https://github.com/raspberrypi-go-drivers/button/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspberrypi-go-drivers/button)](https://goreportcard.com/report/github.com/raspberrypi-go-drivers/button)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This drivers allows interact with a button connected to the GPIO

## Documentation

For full documentation, please visit [![Go Reference](https://pkg.go.dev/badge/github.com/raspberrypi-go-drivers/button.svg)](https://pkg.go.dev/github.com/raspberrypi-go-drivers/button)

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

## Raspberry Pi compatibility

This driver has has only been tested on an Raspberry Pi Zero WH but should work well on every Raspberry Pi

## License

MIT License

---

Special thanks to @stianeikeland

This driver is based on his work in [stianeikeland/go-rpio](https://github.com/stianeikeland/go-rpio/)
