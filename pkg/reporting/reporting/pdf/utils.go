package pdf

import (
	"fmt"
	"github.com/signintech/gopdf"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
	"nokowebapi/nokocore"
	"path/filepath"
	"strings"
	"time"
)

func GetOutputNameForDocPdf(config *Config, index int, date time.Time) string {
	dateFormat := "2006-01-02-15-04-05"
	dateExtend := nokocore.Int64ToBase64RawURL(date.UnixMilli())
	dateOutput := fmt.Sprintf("%s-%s", date.Local().Format(dateFormat), dateExtend)
	indexOutput := fmt.Sprintf("%d", index)
	outputName := strings.ReplaceAll(config.OutputName, "{index}", indexOutput)
	outputName = strings.ReplaceAll(outputName, "{date}", dateOutput)
	return filepath.Join(config.OutputDir, outputName)
}

func GetTemplatePageSize(templateConfig TemplateConfig) gopdf.Rect {
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

func IsTemplateLayoutLandscape(templateConfig TemplateConfig) bool {
	switch strings.ToUpper(templateConfig.PageLayout) {
	case "LANDSCAPE":
		return true

	default:
		return false
	}
}

func IsTemplateLayoutPortrait(templateConfig TemplateConfig) bool {
	switch strings.ToUpper(templateConfig.PageLayout) {
	case "PORTRAIT":
		return true

	default:
		return false
	}
}

func CurrencyFormat(s string, value any) string {
	//lang := language.MustParse(s)
	//cur, unit := currency.FromTag(lang)
	//nokocore.KeepVoid(lang, cur, unit)
	cur := currency.MustParseISO(s)
	scale, inc := currency.Cash.Rounding(cur)
	nokocore.KeepVoid(scale, inc)
	dec := number.Decimal(value, number.Scale(scale))
	//p := message.NewPrinter(lang)
	p := message.NewPrinter(language.English)
	//n := display.Tags(language.English)
	//p.Printf("%24v (%v): %v%v\n", n.Name(lang), cur, currency.Symbol(cur), dec)
	return p.Sprintf("%v", dec)
}
