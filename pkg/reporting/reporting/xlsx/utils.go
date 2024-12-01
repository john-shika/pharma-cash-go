package xlsx

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func GetOutputNameForDocXlsx(config *XlsxConfig, index int, date time.Time) string {
	dateFormat := "2006-01-02-15-04-05"
	outputName := strings.ReplaceAll(config.OutputName, "{index}", fmt.Sprintf("%d", index))
	outputName = strings.ReplaceAll(outputName, "{date}", date.Local().Format(dateFormat))
	return filepath.Join(config.OutputDir, outputName)
}
