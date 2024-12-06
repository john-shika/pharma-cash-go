package hwd

import (
	"errors"
	"fmt"
	"go.bug.st/serial/enumerator"
	"nokowebapi/nokocore"
)

func GetDevicePorts() ([]*enumerator.PortDetails, error) {
	var err error
	var ports []*enumerator.PortDetails

	// maximum qrcode size is ~2968 bytes
	if ports, err = enumerator.GetDetailedPortsList(); err != nil {
		return nil, fmt.Errorf("failed to get serial ports, %w", err)
	}

	size := len(ports)
	if size > 0 {
		for i, port := range ports {
			nokocore.KeepVoid(i)

			fmt.Printf("Found port: %s\n", port.Name)
			if port.IsUSB {
				ID := fmt.Sprintf("%s:%s", port.VID, port.PID)
				fmt.Printf("  USB ID      %s\n", ID)
				fmt.Printf("  USB serial  %s\n", port.SerialNumber)
				fmt.Printf("  USB product %s\n", port.Product)
			}
		}

		return ports, nil
	}

	return nil, errors.New("no serial ports found")
}
