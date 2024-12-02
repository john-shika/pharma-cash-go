package reporting

import (
	"github.com/shopspring/decimal"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"reporting/reporting/pdf"
	"reporting/reporting/xlsx"
)

func PdfService(config *pdf.Config) {
	nokocore.KeepVoid(config)

	doc := pdf.NewDocPdf(1, config)

	formData := &pdf.FormDataPdf{
		Title: "Pharma Cash App",
		Name:  "John, Doe",
		Role:  "Administrator",
		Date:  nokocore.GetTimeUtcNow(),
	}

	rows1 := pdf.Page1TableRowsPdf{}

	for i := range 20 {
		nokocore.KeepVoid(i)
		rows1 = append(rows1, pdf.Page1TableColsPdf{
			Name:   "Paracetamol",
			Buy:    decimal.RequireFromString("10000.00"),
			Margin: decimal.RequireFromString("2000.00"),
			Tax:    decimal.RequireFromString("1200.00"),
			Sale:   decimal.RequireFromString("13200.00"),
			Stock:  10,
			Sold:   5,
			Date:   nokocore.GetTimeUtcNow(),
		})
	}

	doc.Page1Print(formData, rows1)

	rows2 := pdf.Page2TableRowsPdf{}

	for i := range 20 {
		nokocore.KeepVoid(i)
		rows2 = append(rows2, pdf.Page2TableColsPdf{
			Name:     "Paracetamol",
			Officer:  "Angelina, Rose John Doe Jr.",
			Shift:    "Night",
			Quantity: 5,
			Subtotal: decimal.RequireFromString("10000.00"),
			Total:    decimal.RequireFromString("24000000.00"),
			Income:   decimal.RequireFromString("12000000.00"),
			Date:     nokocore.GetTimeUtcNow(),
		})
	}

	doc.Page2Print(formData, rows2)

	rows3 := pdf.Page3TableRowsPdf{}

	for i := range 20 {
		nokocore.KeepVoid(i)
		rows3 = append(rows3, pdf.Page3TableColsPdf{
			Name:     "Paracetamol",
			Brand:    "SoComeOut",
			Supplier: "PT. SoComeOut LTD",
			Stock:    10,
			Expires:  nokocore.GetTimeUtcNow(),
		})
	}

	doc.Page3Print(formData, rows3)

	doc.Save()
}

func XlsxService(config *xlsx.Config) {
	doc := xlsx.NewDocXlsx(0, config)

	formData := &xlsx.FormDataXlsx{
		Title: "Pharma Cash App",
		Name:  "John, Doe",
		Role:  "Administrator",
		Date:  nokocore.GetTimeUtcNow(),
	}

	rows1 := xlsx.Sheet1TableRowsXlsx{
		xlsx.Sheet1TableColsXlsx{
			Name:   "Dragon Fruit",
			Buy:    decimal.RequireFromString("10000.00"),
			Margin: decimal.RequireFromString("2000.00"),
			Tax:    decimal.RequireFromString("1200.00"),
			Sale:   decimal.RequireFromString("13200.00"),
			Stock:  10,
			Sold:   5,
			Date:   nokocore.GetTimeUtcNow(),
		},
	}

	doc.Sheet1Print("Sheet1", formData, rows1)

	rows2 := xlsx.Sheet2TableRowsXlsx{
		{
			Name:     "Paracetamol",
			Brand:    "SoComeOut",
			Supplier: "PT. SoComeOut LTD",
			Stock:    10,
			Expires:  nokocore.GetTimeUtcNow(),
		},
	}

	doc.Sheet2Print("Sheet2", formData, rows2)

	rows3 := xlsx.Sheet3TableRowsXlsx{
		{
			Name:     "Paracetamol",
			Officer:  "Angelina, Rose John Doe Jr.",
			Shift:    "Night",
			Quantity: 5,
			Subtotal: decimal.RequireFromString("10000.00"),
			Total:    decimal.RequireFromString("24000000.00"),
			Income:   decimal.RequireFromString("12000000.00"),
			Date:     nokocore.GetTimeUtcNow(),
		},
	}

	doc.Sheet3Print("Sheet3", formData, rows3)

	doc.Save()
}

func NewService() {
	config := globals.GetConfigGlobals[Config]()
	//PdfService(&config.Pdf)
	XlsxService(&config.Xlsx)
}
