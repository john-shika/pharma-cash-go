package nokocore

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func ToTitleCase(value string) string {
	if IsNoneOrEmptyWhiteSpace(value) {
		return ""
	}

	temp := RegexReplaceAllString("([a-z])([A-Z])", strings.TrimSpace(value), "$1 $2")
	temp = RegexReplaceAllString("([a-zA-Z])([0-9])", temp, "$1 $2")
	temp = RegexReplaceAllString("([0-9])([a-zA-Z])", temp, "$1 $2")
	temp = RegexReplaceAllString("[-_\\s]+", temp, " ")

	//temp = RegexReplaceAllString("\\b\\w", temp, strings.ToUpper("$0"))
	//return temp

	transform := cases.Title(language.English)
	return transform.String(strings.ToLower(temp))
}

func ToUpperStart(value string) string {
	size := len(value)
	switch size {
	case 0:
		return ""
	case 1:
		return strings.ToUpper(value)
	default:
		return strings.ToUpper(value[:1]) + value[1:]
	}
}

func ToLowerStart(value string) string {
	size := len(value)
	switch size {
	case 0:
		return ""
	case 1:
		return strings.ToLower(value)
	default:
		return strings.ToLower(value[:1]) + value[1:]
	}
}

func ToPascalCase(value string) string {
	temp := ToTitleCase(value)
	temp = strings.ReplaceAll(temp, " ", "")
	return ToUpperStart(temp)
}

func ToCamelCase(value string) string {
	temp := ToTitleCase(value)
	temp = strings.ReplaceAll(temp, " ", "")
	return ToLowerStart(temp)
}

func ToSnakeCaseRaw(value string) string {
	temp := ToTitleCase(value)
	temp = strings.ReplaceAll(strings.ToLower(temp), " ", "_")
	return temp
}

func ToSnakeCase(value string) string {
	temp := ToSnakeCaseRaw(value)
	return strings.ToLower(temp)
}

func ToSnakeCaseUpper(value string) string {
	temp := ToSnakeCaseRaw(value)
	return strings.ToUpper(temp)
}

func ToKebabCaseRaw(value string) string {
	temp := ToTitleCase(value)
	temp = strings.ReplaceAll(strings.ToLower(temp), " ", "-")
	return temp
}

func ToKebabCase(value string) string {
	temp := ToKebabCaseRaw(value)
	return strings.ToLower(temp)
}

func ToKebabCaseUpper(value string) string {
	temp := ToKebabCaseRaw(value)
	return strings.ToUpper(temp)
}
