package reporting

import (
	"github.com/signintech/gopdf"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
	"nokowebapi/nokocore"
	"strings"
)

func GetTemplatePageSize(templateConfig PdfTemplateConfig) gopdf.Rect {
	switch strings.ToUpper(templateConfig.PageSize) {
	case "LETTER":
		return gopdf.Rect{W: 612, H: 792}

	case "LEGAL":
		return gopdf.Rect{W: 612, H: 1008}

	case "A1":
		return gopdf.Rect{W: 1685, H: 2384}

	case "A2":
		return gopdf.Rect{W: 1190, H: 1684}

	case "A3":
		return gopdf.Rect{W: 842, H: 1190}

	case "A4":
		return gopdf.Rect{W: 595, H: 842}

	case "A5":
		return gopdf.Rect{W: 420, H: 595}

	case "B4":
		return gopdf.Rect{W: 729, H: 1032}

	case "B5":
		return gopdf.Rect{W: 516, H: 729}

	case "FOLIO":
		return gopdf.Rect{W: 612, H: 936}

	default:
		return gopdf.Rect{W: 0, H: 0}
	}
}

func IsTemplateLayoutLandscape(templateConfig PdfTemplateConfig) bool {
	switch strings.ToUpper(templateConfig.PageLayout) {
	case "LANDSCAPE":
		return true

	default:
		return false
	}
}

func IsTemplateLayoutPortrait(templateConfig PdfTemplateConfig) bool {
	switch strings.ToUpper(templateConfig.PageLayout) {
	case "PORTRAIT":
		return true

	default:
		return false
	}
}

type ReportPdf struct {
	*gopdf.GoPdf
	TemplateConfig PdfTemplateConfig
	Config         PdfConfig
	PageSize       gopdf.Rect
}

func NewReportPdf(config PdfConfig, templateIndex int) *ReportPdf {
	templateConfig := config.Templates[templateIndex]
	pageSize := GetTemplatePageSize(templateConfig)
	if IsTemplateLayoutLandscape(templateConfig) {
		pageSize.W, pageSize.H = pageSize.H, pageSize.W
	}
	goPdf := &gopdf.GoPdf{}
	goPdf.Start(gopdf.Config{
		Unit:              0,
		ConversionForUnit: 0,
		TrimBox:           gopdf.Box{},
		PageSize:          pageSize,
		K:                 0,
		Protection:        gopdf.PDFProtectionConfig{},
	})
	return &ReportPdf{
		GoPdf:          goPdf,
		TemplateConfig: templateConfig,
		Config:         config,
		PageSize:       pageSize,
	}
}

func CurrencyFormat(s string, value any) string {
	//lang := language.MustParse(s)
	//cur, unit := currency.FromTag(lang)
	//nokocore.KeepVoid(cur, unit)
	cur := currency.MustParseISO(s)
	scale, inc := currency.Cash.Rounding(cur)
	nokocore.KeepVoid(scale, inc)
	dec := number.Decimal(value, number.Scale(scale))
	//p := message.NewPrinter(lang)
	p := message.NewPrinter(language.English)
	//n := display.Tags(language.English)
	//p.Printf("%24v (%v): %v%v\n", n.Name(lang), cur, currency.Symbol(cur), dec)
	//return p.Sprintf("%v %v", currency.Symbol(cur), dec)
	return p.Sprintf("%v", dec)
}