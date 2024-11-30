package reporting

import (
	"fmt"
	"github.com/signintech/gopdf"
	"github.com/tealeg/xlsx"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"path/filepath"
	"strings"
)

func PdfService(config *PdfConfig) {
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
	timeUtcNow := nokocore.GetTimeUtcNow()
	dateFormat := "2006-01-02-15-04-05"

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
	pdf.Cell(nil, fmt.Sprintf("%-9s: %s", "Name", "Widho, Hakim"))

	Y += 14
	pdf.SetXY(X, Y)
	pdf.Cell(nil, fmt.Sprintf("%-9s: %s", "Username", "admin"))

	Y += 14
	pdf.SetXY(X, Y)
	pdf.Cell(nil, fmt.Sprintf("%-9s: %s", "Role", "Administrator"))

	Y += 14
	pdf.SetXY(X, Y)
	pdf.Cell(nil, fmt.Sprintf("%-9s: %s", "Date", nokocore.GetTimeUtcNow().Local().Format(nokocore.DateTimeFormat)))

	X = 35
	Y += 20
	pdf.SetXY(X, Y)

	//pdf.Cell(nil, "========================================================================================")

	// Size: 128
	pdf.Cell(nil, "================================================================================================================================")

	Y += 12
	pdf.SetXY(X, Y)

	err = pdf.SetFont(fontFamily, "", fontSize)
	if err != nil {
		console.Error(fmt.Sprintf("panic: %s", err.Error()))
		return
	}

	table := pdf.NewTableLayout(X, Y, 24, 4)
	table.AddColumn("No.", 52, "center")
	table.AddColumn("Name", 224, "left")
	table.AddColumn("Buy Price", 92, "right")
	table.AddColumn("Margin", 64, "right")
	table.AddColumn("TAX", 64, "right")
	table.AddColumn("Sale Price", 92, "right")
	table.AddColumn("Stock", 92, "right")
	table.AddColumn("Date", 92, "right")

	table.AddRow([]string{"1", fmt.Sprintf(" %s", "Dragon Fruit"), CurrencyFormat("IDR", 12000), CurrencyFormat("IDR", 12000), CurrencyFormat("IDR", 12000), CurrencyFormat("IDR", 12000), "32", "2022-10-12 15:45:56 "})
	table.AddRow([]string{"2", fmt.Sprintf(" %s", "Dragon Fruit"), CurrencyFormat("IDR", 12000), CurrencyFormat("IDR", 12000), CurrencyFormat("IDR", 12000), CurrencyFormat("IDR", 12000), "32", "2022-10-12 15:45:56 "})
	table.AddRow([]string{"3", fmt.Sprintf(" %s", "Dragon Fruit"), CurrencyFormat("IDR", 12000), CurrencyFormat("IDR", 12000), CurrencyFormat("IDR", 12000), CurrencyFormat("IDR", 12000), "32", "2022-10-12 15:45:56 "})

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

func XlsxService(config *XlsxConfig) {
	var err error
	var file *xlsx.File

	nokocore.KeepVoid(config)

	fmt.Println(nokocore.ShikaYamlEncode(config))

	index := 0
	template := config.Templates[index]
	timeUtcNow := nokocore.GetTimeUtcNow()
	dateFormat := "2006-01-02-15-04-05"

	outputName := strings.ReplaceAll(config.OutputName, "{index}", fmt.Sprintf("%d", index))
	outputName = strings.ReplaceAll(outputName, "{date}", timeUtcNow.Local().Format(dateFormat))
	outputFilePath := filepath.Join(config.OutputDir, outputName)
	nokocore.KeepVoid(outputFilePath)

	sheetFilePath := filepath.Join(config.Assets, template.SheetFile)
	fmt.Println(sheetFilePath)

	if file, err = xlsx.OpenFile(sheetFilePath); err != nil {
		panic(fmt.Errorf("failed to open file, %w", err))
	}

	format := "_-[$Rp-en-ID]* #,##0.00_-;-[$Rp-en-ID]* #,##0.00_-;_-[$Rp-en-ID]* \"-\"??_-;_-@_-"
	sheet := file.Sheet[template.SheetName]

	style := xlsx.NewStyle()
	style.Border = xlsx.Border{
		Top:    "thin",
		Bottom: "thin",
		Right:  "thin",
		Left:   "thin",
	}

	style.Font = xlsx.Font{
		Size: 11,
		Name: "Arial",
	}

	var cell *xlsx.Cell

	cell = sheet.Cell(1, 2)
	cell.SetFloatWithFormat(10000, format)
	cell.SetStyle(style)

	cell = sheet.Cell(1, 3)
	cell.SetFloatWithFormat(2000, format)
	cell.SetStyle(style)

	cell = sheet.Cell(1, 4)
	cell.SetFloatWithFormat(1200, format)
	cell.SetStyle(style)

	cell = sheet.Cell(1, 5)
	cell.SetFloatWithFormat(13200, format)
	cell.SetStyle(style)

	cell = sheet.Cell(1, 6)
	cell.SetInt(32)
	cell.SetStyle(style)

	file.Save(outputFilePath)
	fmt.Println(outputFilePath)
}

func NewService() {
	config := globals.GetConfigGlobals[Config]()
	XlsxService(&config.Xlsx)
	//PdfService(&config.Pdf)
}
