package funcmap

import (
	"reflect"
	"strings"

	"github.com/twpayne/go-jsonstruct"
)

func ParseFieldName(raw string) string {
	a := &jsonstruct.AbbreviationHandlingFieldNamer{}
	return a.FieldName(raw)
}
func ParseJSONPath(s string) []string {
	return strings.Split(s, ".")
}

func GetFieldType1(i interface{}, s string) string {
	b := reflect.ValueOf(i).Elem().FieldByName(ParseFieldName(s))
	return b.Type().Name()
}

func ParseFieldValue(s string) string {
	parts := strings.Split(s, ".")
	result := ""

	for _, part := range parts {
		if result != "" {
			result += "."
		}
		result += ParseFieldName(part)
	}
	return result

}
