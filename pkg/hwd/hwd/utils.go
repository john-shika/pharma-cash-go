package hwd

import (
	"errors"
	"fmt"
	"go.bug.st/serial/enumerator"
	"log"
	"nokowebapi/console"
	"nokowebapi/nokocore"
)

func GetDevicePorts() ([]*enumerator.PortDetails, error) {
	var err error
	var ports []*enumerator.PortDetails

	// maximum qrcode size is ~2968 bytes
	if ports, err = enumerator.GetDetailedPortsList(); err != nil {
		log.Fatal(err)
	}

	for i, port := range ports {
		nokocore.KeepVoid(i)

		console.Log(fmt.Sprintf("Found port: %s", port.Name))
	}

	if len(ports) == 0 {
		fmt.Println("No serial ports found!")
		return nil, errors.New("no serial ports found")
	}

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
