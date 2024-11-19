package beta

import (
	"context"
	"fmt"
	"github.com/karalabe/hid"
	"github.com/mect/go-escpos"
	"golang.org/x/image/bmp"
	"nokowebapi/nokocore"
	"os"
)

func SessionScanner() {
	for i, device := range hid.Enumerate(0, 0) {
		nokocore.KeepVoid(i)

		fmt.Println(device.Product, device.ProductID, device.VendorID, device.Serial)
	}

	device := hid.Device{
		DeviceInfo: hid.DeviceInfo{
			VendorID:  0x26F1,
			ProductID: 0x5650,
			Serial:    "COM4",
		},
	}

	scanner := nokocore.Unwrap(NewScanner(device.DeviceInfo))

	ctx := context.Background()

	for {
		select {
		case code, ok := <-scanner.ReadCodes(ctx):
			if !ok {
				return
			}

			fmt.Println(code)
		case <-ctx.Done():
			return
		}
	}
}

func SessionPrinter() {
	printer := NewPrinter("PANDA ESCPOS")
	p, err := escpos.NewPrinterByRW(printer)
	if err != nil {
		fmt.Println(err)
		return
	}

	stream, err := os.Open("sales.bmp")
	if err != nil {
		fmt.Println(err)
		return
	}

	//img, format, err := image.Decode(stream)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//fmt.Printf("Format: %s\n", format)

	img, err := bmp.Decode(stream)
	if err != nil {
		fmt.Println(err)
		return
	}

	p.Init()
	p.Smooth(true)

	p.Image(img)

	p.AztecViaImage("HELLO, WORLD, IM FREEDOM!", 100, 100)

	p.Print("\x1dk\x0441234")

	p.Barcode("0123456789", escpos.BarcodeTypeCODE128)

	p.Feed(4)

	p.Cut()
	p.End()

	nokocore.NoErr(printer.Close())
}
