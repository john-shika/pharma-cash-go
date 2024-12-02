package pdf

import (
	"fmt"
	"github.com/signintech/gopdf"
	"nokowebapi/nokocore"
	"strings"
)

var dateTimeFormat = "Monday, 2 January 2006"

func SetFormDataPdf(pdf *gopdf.GoPdf, X float64, Y float64, W float64, H float64, formData *FormDataPdf) (err error) {
	nokocore.KeepVoid(pdf, X, Y, W, H, formData)

	date := formData.Date.Local().Format(dateTimeFormat)
	size := len(formData.Title) * 12
	m := size / 2
	k := (W / 2) - float64(m)
	nokocore.KeepVoid(m, k)

	pdf.SetXY(k, 52)
	if err = pdf.SetFontSize(16); err != nil {
		return err
	}

	if err = pdf.Cell(nil, formData.Title); err != nil {
		return err
	}

	pdf.SetXY(X, Y)
	if err = pdf.SetFontSize(12); err != nil {
		return err
	}

	if err = pdf.Cell(nil, fmt.Sprintf("%-6s: %s", "Name", formData.Name)); err != nil {
		return err
	}

	Y += 14
	pdf.SetXY(X, Y)
	if err = pdf.Cell(nil, fmt.Sprintf("%-6s: %s", "Role", formData.Role)); err != nil {
		return err
	}

	Y += 14
	pdf.SetXY(X, Y)
	if err = pdf.Cell(nil, fmt.Sprintf("%-6s: %s", "Date", date)); err != nil {
		return err
	}

	X = 35
	Y += 34
	pdf.SetXY(X, Y)
	if err = pdf.Cell(nil, strings.Repeat("=", 128)); err != nil {
		return err
	}

	return nil
}
