package utils

import (
	"fmt"
	"nokowebapi/nokocore"
	"strings"
)

func ToShiftNameNorm(name string) string {
	name = nokocore.ToTitleCase(name)
	if !strings.HasSuffix(name, "Shift") {
		return fmt.Sprintf("%s Shift", name)
	}
	return name
}
