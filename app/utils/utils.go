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

func Modulo(value, m int) (extra int, div int) {
	div = 0
	extra = value
	for extra >= m {
		extra += -m
		div += 1
	}
	return extra, div
}
