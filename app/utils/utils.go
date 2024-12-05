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

	switch {
	case value > 0 && m > 0:
		for extra >= m {
			extra += -m
			div += 1
		}

	case value < 0 && m > 0:
		extra, div = Modulo(-value, m)
		extra = m - extra

	case value > 0 && m < 0:
		extra, div = Modulo(value, -m)
		extra = m + extra

	case value < 0 && m < 0:
		extra, div = Modulo(-value, -m)
		extra = -extra

	default:
		panic("division by zero")
	}

	return extra, div
}
