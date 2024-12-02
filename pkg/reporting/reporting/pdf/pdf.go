package pdf

import (
	"fmt"
	"github.com/signintech/gopdf"
	"nokowebapi/nokocore"
	"path/filepath"
	"strconv"
	"time"
)

type DocPdfImpl interface {
	SlicePage1Print(formData *FormDataPdf, start int, rows Page1TableRowsPdf) (err error)
	SlicePage2Print(formData *FormDataPdf, start int, rows Page2TableRowsPdf) (err error)
	SlicePage3Print(formData *FormDataPdf, start int, rows Page3TableRowsPdf) (err error)
	Page1Print(formData *FormDataPdf, rows Page1TableRowsPdf) (err error)
	Page2Print(formData *FormDataPdf, rows Page2TableRowsPdf) (err error)
	Page3Print(formData *FormDataPdf, rows Page3TableRowsPdf) (err error)
	Save() (err error)
}

type DocPdf struct {
	PDF            *gopdf.GoPdf
	Config         *Config
	TemplatePageId int
	TableStyle     gopdf.CellStyle
	PageSize       gopdf.Rect
	Date           time.Time
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
			R: 0,
			G: 0,
			B: 0,
		},
		Font:     fontFamily,
		FontSize: config.FontSize,
	}

	templateFilePath := filepath.Join(config.Assets, template.PageFile)
	templatePageId := pdf.ImportPage(templateFilePath, 1, "/MediaBox")

	return &DocPdf{
		PDF:            pdf,
		Config:         config,
		TemplatePageId: templatePageId,
		TableStyle:     tableStyle,
		PageSize:       pageSize,
		Date:           time.Now(),
	}
}

func (d *DocPdf) SlicePage1Print(formData *FormDataPdf, start int, rows Page1TableRowsPdf) (err error) {
	var X, Y float64

	d.PDF.AddPage()
	d.PDF.UseImportedTemplate(d.TemplatePageId, 0, 0, d.PageSize.W, d.PageSize.H)

	X = d.PageSize.W - 240
	Y = 38

	W, H := d.PageSize.W, d.PageSize.H
	if err = SetFormDataPdf(d.PDF, X, Y, W, H, formData); err != nil {
		return err
	}

	X = 35
	Y = 112

	d.PDF.SetXY(X, Y)
	if err = d.PDF.SetFontSize(d.Config.FontSize); err != nil {
		return err
	}

	maxRows := len(rows)
	table := d.PDF.NewTableLayout(X, Y, 24, maxRows)
	table.AddColumn("No", 40, "center")
	table.AddColumn("Name", 128, "left")
	table.AddColumn("Buy", 80, "right")
	table.AddColumn("Margin", 80, "right")
	table.AddColumn("TAX", 80, "right")
	table.AddColumn("Sale", 80, "right")
	table.AddColumn("Stock", 64, "right")
	table.AddColumn("Sold", 64, "right")
	table.AddColumn("Date", 156, "right")

	for i, row := range rows {
		nokocore.KeepVoid(i)
		j := start + i + 1

		table.AddRow([]string{
			strconv.Itoa(j),
			nokocore.LimitTextContent(row.Name, 20),
			nokocore.DecimalCurrencyFormat(row.Buy),
			nokocore.DecimalCurrencyFormat(row.Margin),
			nokocore.DecimalCurrencyFormat(row.Tax),
			nokocore.DecimalCurrencyFormat(row.Sale),
			strconv.Itoa(row.Stock),
			strconv.Itoa(row.Sold),
			row.Date.Format(dateTimeFormat),
		})
	}

	table.SetTableStyle(d.TableStyle)
	if err = table.DrawTable(); err != nil {
		return err
	}

	return nil
}

func (d *DocPdf) SlicePage2Print(formData *FormDataPdf, start int, rows Page2TableRowsPdf) (err error) {
	var X, Y float64

	d.PDF.AddPage()
	d.PDF.UseImportedTemplate(d.TemplatePageId, 0, 0, d.PageSize.W, d.PageSize.H)

	X = d.PageSize.W - 240
	Y = 38

	W, H := d.PageSize.W, d.PageSize.H
	if err = SetFormDataPdf(d.PDF, X, Y, W, H, formData); err != nil {
		return err
	}

	X = 35
	Y = 112

	d.PDF.SetXY(X, Y)
	if err = d.PDF.SetFontSize(d.Config.FontSize); err != nil {
		return err
	}

	maxRows := len(rows)
	table := d.PDF.NewTableLayout(X, Y, 24, maxRows)
	table.AddColumn("No", 40, "center")
	table.AddColumn("Name", 128, "left")
	table.AddColumn("Officer", 128, "left")
	table.AddColumn("Shift", 40, "center")
	table.AddColumn("Quantity", 64, "right")
	table.AddColumn("Subtotal", 72, "right")
	table.AddColumn("Total", 72, "right")
	table.AddColumn("Income", 72, "right")
	table.AddColumn("Date", 156, "right")

	for i, row := range rows {
		nokocore.KeepVoid(i)
		j := start + i + 1

		table.AddRow([]string{
			strconv.Itoa(j),
			nokocore.LimitTextContent(row.Name, 20),
			nokocore.LimitTextContent(row.Officer, 20),
			nokocore.LimitTextContent(row.Shift, 5),
			strconv.Itoa(row.Quantity),
			nokocore.DecimalCurrencyFormat(row.Subtotal),
			nokocore.DecimalCurrencyFormat(row.Total),
			nokocore.DecimalCurrencyFormat(row.Income),
			row.Date.Format(dateTimeFormat),
		})
	}

	table.SetTableStyle(d.TableStyle)
	if err = table.DrawTable(); err != nil {
		return err
	}

	return nil
}

func (d *DocPdf) SlicePage3Print(formData *FormDataPdf, start int, rows Page3TableRowsPdf) (err error) {
	var X, Y float64

	d.PDF.AddPage()
	d.PDF.UseImportedTemplate(d.TemplatePageId, 0, 0, d.PageSize.W, d.PageSize.H)

	X = d.PageSize.W - 240
	Y = 38

	W, H := d.PageSize.W, d.PageSize.H
	if err = SetFormDataPdf(d.PDF, X, Y, W, H, formData); err != nil {
		return err
	}

	X = 35
	Y = 112

	d.PDF.SetXY(X, Y)
	if err = d.PDF.SetFontSize(d.Config.FontSize); err != nil {
		return err
	}

	maxRows := len(rows)
	table := d.PDF.NewTableLayout(X, Y, 24, maxRows)
	table.AddColumn("No", 40, "center")
	table.AddColumn("Name", 128, "left")
	table.AddColumn("Brand", 128, "left")
	table.AddColumn("Supplier", 256, "left")
	table.AddColumn("Stock", 64, "right")
	table.AddColumn("Expires", 156, "right")

	for i, row := range rows {
		nokocore.KeepVoid(i)
		j := start + i + 1

		table.AddRow([]string{
			strconv.Itoa(j),
			nokocore.LimitTextContent(row.Name, 20),
			nokocore.LimitTextContent(row.Brand, 20),
			nokocore.LimitTextContent(row.Supplier, 40),
			strconv.Itoa(row.Stock),
			row.Expires.Format(dateTimeFormat),
		})
	}

	table.SetTableStyle(d.TableStyle)
	if err = table.DrawTable(); err != nil {
		return err
	}

	return nil
}

func (d *DocPdf) Page1Print(formData *FormDataPdf, rows Page1TableRowsPdf) (err error) {
	formData2 := &FormDataPdf{
		Title: formData.Title,
		Name:  formData.Name,
		Role:  formData.Role,
		Date:  formData.Date,
	}
	title := formData2.Title
	size := len(rows)
	m := 16
	for i := 0; i < size; i += m {
		j := (i / m) + 1
		formData2.Title = fmt.Sprintf("%s (%d)", title, j)

		x := i + m
		if size > x {
			temp := make(Page1TableRowsPdf, m)
			copy(temp, rows[i:x])
			if err = d.SlicePage1Print(formData2, i, temp); err != nil {
				return err
			}
			continue
		}

		k := size - i
		x = i + k
		temp := make(Page1TableRowsPdf, k)
		copy(temp, rows[i:x])
		if err = d.SlicePage1Print(formData2, i, temp); err != nil {
			return err
		}
	}

	return nil
}

func (d *DocPdf) Page2Print(formData *FormDataPdf, rows Page2TableRowsPdf) (err error) {
	formData2 := &FormDataPdf{
		Title: formData.Title,
		Name:  formData.Name,
		Role:  formData.Role,
		Date:  formData.Date,
	}
	title := formData2.Title
	size := len(rows)
	m := 16
	for i := 0; i < size; i += m {
		j := (i / m) + 1
		formData2.Title = fmt.Sprintf("%s (%d)", title, j)

		x := i + m
		if size > x {
			temp := make(Page2TableRowsPdf, m)
			copy(temp, rows[i:x])
			if err = d.SlicePage2Print(formData2, i, temp); err != nil {
				return err
			}
			continue
		}

		k := size - i
		x = i + k
		temp := make(Page2TableRowsPdf, k)
		copy(temp, rows[i:x])
		if err = d.SlicePage2Print(formData2, i, temp); err != nil {
			return err
		}
	}

	return nil
}

func (d *DocPdf) Page3Print(formData *FormDataPdf, rows Page3TableRowsPdf) (err error) {
	formData2 := &FormDataPdf{
		Title: formData.Title,
		Name:  formData.Name,
		Role:  formData.Role,
		Date:  formData.Date,
	}
	title := formData2.Title
	size := len(rows)
	m := 16
	for i := 0; i < size; i += m {
		j := (i / m) + 1
		formData2.Title = fmt.Sprintf("%s (%d)", title, j)

		x := i + m
		if size > x {
			temp := make(Page3TableRowsPdf, m)
			copy(temp, rows[i:x])
			if err = d.SlicePage3Print(formData2, i, temp); err != nil {
				return err
			}
			continue
		}

		k := size - i
		x = i + k
		temp := make(Page3TableRowsPdf, k)
		copy(temp, rows[i:x])
		if err = d.SlicePage3Print(formData2, i, temp); err != nil {
			return err
		}
	}

	return nil
}

func (d *DocPdf) Save() (err error) {
	outputFilePath := GetOutputNameForDocPdf(d.Config, d.TemplatePageId, d.Date)
	if err = d.PDF.WritePdf(outputFilePath); err != nil {
		return err
	}

	return nil
}
