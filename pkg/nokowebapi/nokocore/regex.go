package nokocore

import (
	"io"
	"regexp"
	"sync"
)

const (
	booleanRegex = "(?i)^(0|1|n(o)?|y(es)?|false|true)$"
	numberRegex  = "^([0-9]+)$"
	numericRegex = "^[\\+\\-]?([0-9]+)(\\.([0-9]+))?$"
	alphaRegex   = "^([a-zA-Z]+)$"
	asciiRegex   = "^([ -~]+)$"
	emailRegex   = "^([a-zA-Z0-9]+)(([\\_\\.\\-]([a-zA-Z0-9]+))+)?\\@([a-zA-Z0-9]+)(([\\_\\.\\-]([a-zA-Z0-9]+))+)?$"
	phoneRegex   = "^(\\+[0-9]{1,2}|0)\\s?([0-9]{3,4}(-|\\s)?){2,3}[0-9]{3,4}$"
)

var (
	BooleanRegex = LazyRegexCompile(booleanRegex)
	NumberRegex  = LazyRegexCompile(numberRegex)
	NumericRegex = LazyRegexCompile(numericRegex)
	AlphaRegex   = LazyRegexCompile(alphaRegex)
	AsciiRegex   = LazyRegexCompile(asciiRegex)
	EmailRegex   = LazyRegexCompile(emailRegex)
	PhoneRegex   = LazyRegexCompile(phoneRegex)
)

type RegexpImpl interface {
	String() string
	Copy() *regexp.Regexp
	Longest()
	NumSubexp() int
	SubexpNames() []string
	SubexpIndex(name string) int
	LiteralPrefix() (prefix string, complete bool)
	MatchReader(r io.RuneReader) bool
	MatchString(s string) bool
	Match(b []byte) bool
	ReplaceAllString(src string, repl string) string
	ReplaceAllLiteralString(src string, repl string) string
	ReplaceAllStringFunc(src string, repl func(string) string) string
	ReplaceAll(src []byte, repl []byte) []byte
	ReplaceAllLiteral(src []byte, repl []byte) []byte
	ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
	Find(b []byte) []byte
	FindIndex(b []byte) (loc []int)
	FindString(s string) string
	FindStringIndex(s string) (loc []int)
	FindReaderIndex(r io.RuneReader) (loc []int)
	FindSubmatch(b []byte) [][]byte
	Expand(dst []byte, template []byte, src []byte, match []int) []byte
	ExpandString(dst []byte, template string, src string, match []int) []byte
	FindSubmatchIndex(b []byte) []int
	FindStringSubmatch(s string) []string
	FindStringSubmatchIndex(s string) []int
	FindReaderSubmatchIndex(r io.RuneReader) []int
	FindAll(b []byte, n int) [][]byte
	FindAllIndex(b []byte, n int) [][]int
	FindAllString(s string, n int) []string
	FindAllStringIndex(s string, n int) [][]int
	FindAllSubmatch(b []byte, n int) [][][]byte
	FindAllSubmatchIndex(b []byte, n int) [][]int
	FindAllStringSubmatch(s string, n int) [][]string
	FindAllStringSubmatchIndex(s string, n int) [][]int
	Split(s string, n int) []string
	MarshalText() ([]byte, error)
	UnmarshalText(text []byte) error
}

var cachesRegex = make(map[string]RegexpImpl)

type ReOrStrImpl interface {
	string | *regexp.Regexp
}

func GetRegexPattern[T ReOrStrImpl](pattern T) RegexpImpl {
	var ok bool
	var re RegexpImpl
	var str string
	KeepVoid(ok, re, str)

	// check pattern is regex pointer or string
	if re, ok = Cast[RegexpImpl](pattern); !ok {
		if str, ok = Cast[string](pattern); !ok {
			panic("pattern must be 'regexp.Regexp' or string type")
		}

		// register new regex and store in cachesRegex
		if re, ok = cachesRegex[str]; !ok {
			re = regexp.MustCompile(str)
			cachesRegex[str] = re
			return re
		}

		return re
	}

	return nil
}

func LazyRegexCompile[T ReOrStrImpl](pattern T) func() RegexpImpl {
	var regex RegexpImpl
	var once sync.Once
	return func() RegexpImpl {
		once.Do(func() {
			regex = GetRegexPattern(pattern)
		})
		return regex
	}
}

func RegexMatch[T ReOrStrImpl](pattern T, value []byte) bool {
	re := GetRegexPattern(pattern)
	return re.Match(value)
}

func RegexMatchString[T ReOrStrImpl](pattern T, value string) bool {
	re := GetRegexPattern(pattern)
	return re.MatchString(value)
}

func RegexFind[T ReOrStrImpl](pattern T, value []byte) []byte {
	re := GetRegexPattern(pattern)
	return re.Find(value)
}

func RegexFindString[T ReOrStrImpl](pattern T, value string) string {
	re := GetRegexPattern(pattern)
	return re.FindString(value)
}

func RegexFindAll[T ReOrStrImpl](pattern T, value []byte, n int) [][]byte {
	re := GetRegexPattern(pattern)
	return re.FindAll(value, n)
}

func RegexFindAllString[T ReOrStrImpl](pattern T, value string, n int) []string {
	re := GetRegexPattern(pattern)
	return re.FindAllString(value, n)
}

func RegexReplaceAll[T ReOrStrImpl](pattern T, value []byte, replace []byte) []byte {
	re := GetRegexPattern(pattern)
	return re.ReplaceAll(value, replace)
}

func RegexReplaceAllString[T ReOrStrImpl](pattern T, value string, replace string) string {
	re := GetRegexPattern(pattern)
	return re.ReplaceAllString(value, replace)
}

func RegexSplit[T ReOrStrImpl](pattern T, value string, n int) []string {
	re := GetRegexPattern(pattern)
	return re.Split(value, n)
}
