package reporting

import (
	"fmt"
	"github.com/signintech/gopdf"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
	"os"
	"path/filepath"
)

func NewService() {
	config := globals.GetConfigGlobals[Config]()
	fmt.Println(nokocore.ShikaYamlEncode(config))

	fontFamily := nokocore.ToPascalCase(config.FontFamily)
	fontFileName := fmt.Sprintf("%s-%s.ttf", fontFamily, config.FontType)
	fontFilePath := filepath.Join(config.Assets, "fonts", fontFamily, fontFileName)

	if file, err := os.OpenFile(fontFilePath, os.O_RDONLY, 0); err != nil {
		nokocore.KeepVoid(file)
		panic(fmt.Errorf("failed to open file, %w", err))
	}

	fmt.Println(fontFilePath)

	pdfFilePath := filepath.Join(config.OutputDir, "example.pdf")

	pdf := gopdf.GoPdf{}

	pageSize := gopdf.Rect{W: 595, H: 842}
	pdf.Start(gopdf.Config{PageSize: pageSize})

	pdf.AddPage()
	err := pdf.AddTTFFont(fontFamily, fontFilePath)
	if err != nil {
		console.Error(fmt.Sprintf("panic: %s", err.Error()))
		return
	}

	var fontSize int

	fontSize = 12
	err = pdf.SetFont(fontFamily, "", fontSize)
	if err != nil {
		console.Error(fmt.Sprintf("panic: %s", err.Error()))
		return
	}

	templateFilePath := filepath.Join(config.Assets, "templates", "template.pdf")
	template1 := pdf.ImportPage(templateFilePath, 1, "/MediaBox")
	pdf.UseImportedTemplate(template1, 0, 0, pageSize.W, pageSize.H)

	//pdf.SetLineWidth(0.1)
	//pdf.SetFillColor(124, 252, 0) //setup fill color
	//pdf.RectFromUpperLeftWithStyle(50, 100, 400, 600, "FD")
	//pdf.SetFillColor(0, 0, 0)

	var X, Y float64

	X = pageSize.W - 240
	Y = 20

	pdf.SetXY(X, Y)
	pdf.Cell(nil, fmt.Sprintf("%-9s: %s", "Name", "Widho, Hakim"))

	Y += 16
	pdf.SetXY(X, Y)
	pdf.Cell(nil, fmt.Sprintf("%-9s: %s", "Username", "admin"))

	Y += 16
	pdf.SetXY(X, Y)
	pdf.Cell(nil, fmt.Sprintf("%-9s: %s", "Role", "Administrator"))

	Y += 16
	pdf.SetXY(X, Y)
	pdf.Cell(nil, fmt.Sprintf("%-9s: %s", "Date", nokocore.GetTimeUtcNow().Local().Format(nokocore.DateTimeFormat)))

	X = 34
	Y += 32
	pdf.SetXY(X, Y)

	// Size: 88
	pdf.Cell(nil, "========================================================================================")

	Y += 12
	pdf.SetXY(X, Y)

	fontSize = 6
	err = pdf.SetFont(fontFamily, "", fontSize)
	if err != nil {
		console.Error(fmt.Sprintf("panic: %s", err.Error()))
		return
	}

	table := pdf.NewTableLayout(X, Y, 12, 4)
	table.AddColumn("No.", 32, "center")
	table.AddColumn("Name", 124, "center")
	table.AddColumn("Buy Price", 64, "center")
	table.AddColumn("Sale Price", 64, "center")
	table.AddColumn("Margin", 42, "center")
	table.AddColumn("TAX", 42, "center")
	table.AddColumn("Total", 80, "center")
	table.AddColumn("Income", 80, "center")
	table.AddRow([]string{"1", "Dragon Fruit", "10000", "13200", "2000", "1200", "13200", "2000"})
	table.AddRow([]string{"2", "Dragon Fruit", "10000", "13200", "2000", "1200", "13200", "2000"})
	table.AddRow([]string{"3", "Dragon Fruit", "10000", "13200", "2000", "1200", "13200", "2000"})
	table.DrawTable()

	pdf.WritePdf(pdfFilePath)
}
