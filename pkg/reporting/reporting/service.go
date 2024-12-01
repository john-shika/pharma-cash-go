package reporting

import (
	"fmt"
	"github.com/signintech/gopdf"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"path/filepath"
	"reporting/reporting/pdf"
	"reporting/reporting/xlsx"
	"strings"
)

func PdfService(config *pdf.Config) {
	nokocore.KeepVoid(config)

	fontFamily := nokocore.ToPascalCase(config.FontFamily)
	fontFileName := fmt.Sprintf("%s-%s.ttf", fontFamily, config.FontType)
	fontFilePath := filepath.Join(config.Assets, "fonts", fontFamily, fontFileName)

	//if file, err := os.OpenFile(fontFilePath, os.O_RDONLY, 0); err != nil {
	//	nokocore.KeepVoid(file)
	//	panic(fmt.Errorf("failed to open file, %w", err))
	//}

	fmt.Println(fontFilePath)

	index := 1
	template := config.Templates[index]

	dateFormat := "2006-01-02-15-04-05"
	timeUtcNow := nokocore.GetTimeUtcNow()
	outputName := strings.ReplaceAll(config.OutputName, "{index}", fmt.Sprintf("%d", index))
	outputName = strings.ReplaceAll(outputName, "{date}", timeUtcNow.Local().Format(dateFormat))
	outputFilePath := filepath.Join(config.OutputDir, outputName)

	pdf := NewReportPdf(*config, index)
	pageSize := pdf.PageSize

	pdf.AddPage()
	err := pdf.AddTTFFont(fontFamily, fontFilePath)
	if err != nil {
		console.Error(fmt.Sprintf("panic: %s", err.Error()))
		return
	}

	fontSize := config.FontSize
	err = pdf.SetFont(fontFamily, "", fontSize)
	if err != nil {
		console.Error(fmt.Sprintf("panic: %s", err.Error()))
		return
	}

	templateFilePath := filepath.Join(config.Assets, template.PageFile)
	template1 := pdf.ImportPage(templateFilePath, 1, "/MediaBox")
	pdf.UseImportedTemplate(template1, 0, 0, pageSize.W, pageSize.H)

	//GoPdf.SetLineWidth(0.1)
	//GoPdf.SetFillColor(124, 252, 0) //setup fill color
	//GoPdf.RectFromUpperLeftWithStyle(50, 100, 400, 600, "FD")
	//GoPdf.SetFillColor(0, 0, 0)

	var X, Y float64

	X = pageSize.W - 240
	Y = 38

	pdf.SetXY(X, Y)
	pdf.Cell(nil, fmt.Sprintf("%-6s: %s", "Name", "Widho, Hakim"))

	Y += 14
	pdf.SetXY(X, Y)
	pdf.Cell(nil, fmt.Sprintf("%-6s: %s", "Role", "Administrator"))

	Y += 14
	pdf.SetXY(X, Y)
	pdf.Cell(nil, fmt.Sprintf("%-6s: %s", "Date", nokocore.GetTimeUtcNow().Local().Format(nokocore.DateTimeFormat)))

	X = 35
	Y += 34
	pdf.SetXY(X, Y)

	// Size: 128
	pdf.Cell(nil, strings.Repeat("=", 128))

	Y += 12
	pdf.SetXY(X, Y)

	table := pdf.NewTableLayout(X, Y, 24, 4)
	table.AddColumn("No.", 40, "center")
	table.AddColumn("Name", 156, "left")
	table.AddColumn("Buy", 72, "right")
	table.AddColumn("Margin", 72, "right")
	table.AddColumn("TAX", 72, "right")
	table.AddColumn("Sale", 72, "right")
	table.AddColumn("Stock In", 64, "right")
	table.AddColumn("Stock Out", 64, "right")
	table.AddColumn("Date", 156, "right")

	code := "IDR"
	dateTimeFormat := "Monday, 2 January 2006"
	table.AddRow([]string{"1", fmt.Sprintf("%1s", "Dragon Fruit"), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), "32", "0", timeUtcNow.Format(dateTimeFormat)})
	table.AddRow([]string{"2", fmt.Sprintf("%1s", "Dragon Fruit"), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), "32", "0", timeUtcNow.Format(dateTimeFormat)})
	table.AddRow([]string{"3", fmt.Sprintf("%1s", "Dragon Fruit"), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), "32", "0", timeUtcNow.Format(dateTimeFormat)})

	tableStyle := gopdf.CellStyle{
		BorderStyle: gopdf.BorderStyle{
			Top:    true,
			Left:   true,
			Right:  true,
			Bottom: true,
			Width:  1,
			RGBColor: gopdf.RGBColor{
				R: 0,
				G: 0,
				B: 0,
			},
		},
		FillColor: gopdf.RGBColor{
			R: 255,
			G: 255,
			B: 255,
		},
		TextColor: gopdf.RGBColor{
			R: 255,
			G: 255,
			B: 255,
		},
		Font:     fontFamily,
		FontSize: float64(fontSize),
	}
	table.SetTableStyle(tableStyle)
	table.DrawTable()

	pdf.WritePdf(outputFilePath)
}

func XlsxService(config *xlsx.Config) {
	doc := xlsx.NewDocXlsx(0, config)

	formData := xlsx.FormDataXlsx{
		Name: "Widho, Hakim",
		Role: "Administrator",
		Date: nokocore.GetTimeUtcNow(),
	}
	doc.Sheet1Print("Sheet1", &formData, xlsx.Sheet1TableRowsXlsx{})
	doc.Sheet2Print("Sheet2", &formData, xlsx.Sheet2TableRowsXlsx{})

	doc.Save()
}

func NewService() {
	config := globals.GetConfigGlobals[Config]()
	//PdfService(&config.Pdf)
	XlsxService(&config.Xlsx)
}
