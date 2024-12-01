package pdf

import (
	"fmt"
	"github.com/signintech/gopdf"
	"path/filepath"
	"strings"
	"time"
)

func GetOutputNameForDocPdf(config *PdfConfig, index int, date time.Time) string {
	dateFormat := "2006-01-02-15-04-05"
	outputName := strings.ReplaceAll(config.OutputName, "{index}", fmt.Sprintf("%d", index))
	outputName = strings.ReplaceAll(outputName, "{date}", date.Local().Format(dateFormat))
	return filepath.Join(config.OutputDir, outputName)
}

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
