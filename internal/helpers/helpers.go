package helpers

import (
	"strings"
	"unicode"
)

func RemoveWhitespaces(s string) string {
	words := strings.FieldsFunc(s, unicode.IsSpace)
	return strings.Join(words, "")
}
