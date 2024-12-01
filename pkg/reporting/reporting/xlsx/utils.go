package xlsx

import (
	"fmt"
	"nokowebapi/nokocore"
	"path/filepath"
	"strings"
	"time"
)

func GetOutputNameForDocXlsx(config *Config, index int, date time.Time) string {
	dateFormat := "2006-01-02-15-04-05"
	dateExtend := nokocore.Int64ToBase64RawURL(date.UnixMilli())
	dateOutput := fmt.Sprintf("%s-%s", date.Local().Format(dateFormat), dateExtend)
	indexOutput := fmt.Sprintf("%d", index)
	outputName := strings.ReplaceAll(config.OutputName, "{index}", indexOutput)
	outputName = strings.ReplaceAll(outputName, "{date}", dateOutput)
	return filepath.Join(config.OutputDir, outputName)
}
