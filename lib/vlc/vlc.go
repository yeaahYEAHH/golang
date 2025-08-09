package vlc

import (
	"strings"
	"unicode"
)

func Encode(str string) string {
	// prepare text: M -> !m

	// encode to binary: some text -> 10010101

	return ""
}

func prepareText(str string) string {
	var buf strings.Builder

	for _, char := range str {
		if unicode.IsUpper(char) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(char))
		} else {
			buf.WriteRune(char)
		}
	}

	return buf.String()
}
