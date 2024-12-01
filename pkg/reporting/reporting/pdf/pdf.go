package pdf

import (
	"fmt"
	"github.com/signintech/gopdf"
	"nokowebapi/nokocore"
	"path/filepath"
	"strings"
	"time"
)

type DocPdfImpl interface {
}

type DocPdf struct {
	PDF        *gopdf.GoPdf
	Config     *Config
	TableStyle gopdf.CellStyle
	PageSize   gopdf.Rect
	Page       int
	Date       time.Time
}

func NewDocPdf(templateId int, config *Config) DocPdfImpl {
	fontFamily := nokocore.ToPascalCase(config.FontFamily)
	fontFileName := fmt.Sprintf("%s-%s.ttf", fontFamily, config.FontType)
	fontFilePath := filepath.Join(config.Assets, "fonts", fontFamily, fontFileName)

	template := config.Templates[templateId]
	pageSize := GetTemplatePageSize(template)
	if IsTemplateLayoutLandscape(template) {
		pageSize.W, pageSize.H = pageSize.H, pageSize.W
	}

	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		Unit:              0,
		ConversionForUnit: 0,
		TrimBox:           gopdf.Box{},
		PageSize:          pageSize,
		K:                 0,
		Protection:        gopdf.PDFProtectionConfig{},
	})

	if err := pdf.AddTTFFont(fontFamily, fontFilePath); err != nil {
		panic(err)
	}

	if err := pdf.SetFont(fontFamily, "", config.FontSize); err != nil {
		panic(err)
	}

	templateFilePath := filepath.Join(config.Assets, template.PageFile)
	templatePage1 := pdf.ImportPage(templateFilePath, 1, "/MediaBox")
	pdf.UseImportedTemplate(templatePage1, 0, 0, pageSize.W, pageSize.H)

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
		FontSize: config.FontSize,
	}

	return &DocPdf{
		PDF:        pdf,
		Config:     config,
		TableStyle: tableStyle,
		PageSize:   pageSize,
		Page:       1,
		Date:       time.Now(),
	}
}

func (d *DocPdf) Page1Print(formData *FormDataPdf, rows *Page1TableRowsPdf) (err error) {
	var X, Y float64

	X = d.PageSize.W - 240
	Y = 38

	d.PDF.SetXY(X, Y)
	d.PDF.Cell(nil, fmt.Sprintf("%-6s: %s", "Name", formData.Name))

	Y += 14
	d.PDF.SetXY(X, Y)
	d.PDF.Cell(nil, fmt.Sprintf("%-6s: %s", "Role", formData.Role))

	Y += 14
	d.PDF.SetXY(X, Y)
	d.PDF.Cell(nil, fmt.Sprintf("%-6s: %s", "Date", nokocore.GetTimeUtcNow().Local().Format(nokocore.DateTimeFormat)))

	X = 35
	Y += 34
	d.PDF.SetXY(X, Y)

	// Size: 128
	d.PDF.Cell(nil, strings.Repeat("=", 128))

	Y += 12
	d.PDF.SetXY(X, Y)

	table := d.PDF.NewTableLayout(X, Y, 24, 4)
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
	timeUtcNow := nokocore.GetTimeUtcNow()
	table.AddRow([]string{"1", fmt.Sprintf("%1s", "Dragon Fruit"), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), "32", "0", timeUtcNow.Format(dateTimeFormat)})
	table.AddRow([]string{"2", fmt.Sprintf("%1s", "Dragon Fruit"), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), "32", "0", timeUtcNow.Format(dateTimeFormat)})
	table.AddRow([]string{"3", fmt.Sprintf("%1s", "Dragon Fruit"), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), CurrencyFormat(code, 12000), "32", "0", timeUtcNow.Format(dateTimeFormat)})

	table.SetTableStyle(d.TableStyle)
	table.DrawTable()

	return nil
}
