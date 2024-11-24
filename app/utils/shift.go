package utils

import (
	"fmt"
	"nokowebapi/nokocore"
	"strings"
)

func ToShiftName(name string) string {
	name = nokocore.ToTitleCase(strings.TrimSpace(name))
	if !strings.HasSuffix(name, "Shift") {
		return fmt.Sprintf("%s Shift", name)
	}
	return name
}
