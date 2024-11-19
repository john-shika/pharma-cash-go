package beta

import (
	"context"
	"fmt"
	"github.com/karalabe/hid"
	"nokowebapi/nokocore"
)

// Scanner describes the symbol scanner device
type Scanner struct {
	device *hid.Device
}

// NewScanner takes a HID device info and returns a Scanner with the opened device
func NewScanner(deviceInfo hid.DeviceInfo) (*Scanner, error) {

	device, err := deviceInfo.Open()
	if err != nil {
		return nil, err
	}

	return &Scanner{
		device: device,
	}, nil
}

func (s *Scanner) Read(p []byte) (int, error) {
	return s.device.Read(p)
}

func (s *Scanner) Write(p []byte) (int, error) {
	return s.device.Write(p)
}

// ReadCodes starts a code read loop and returns a channel that receives new codes
func (s *Scanner) ReadCodes(ctx context.Context) <-chan *Code {

	scanCtx, cancel := context.WithCancel(ctx)

	codeChan := make(chan *Code, 1)
	go func() {
		defer close(codeChan)
		for {
			buf := make([]byte, 255)
			n, err := s.device.Read(buf)
			if err != nil {
				cancel()
				return
			}

			if n > 0 {
				fmt.Println(buf)
				scannedCode, err := NewCode(buf)
				if err != nil {
					continue
				}

				codeChan <- scannedCode
			}

			select {
			case <-scanCtx.Done():
				nokocore.NoErr(s.device.Close())
				return
			default:
			}
		}
	}()

	return codeChan
}

// Device returns the real device
func (s *Scanner) Device() *hid.Device {
	return s.device
}
