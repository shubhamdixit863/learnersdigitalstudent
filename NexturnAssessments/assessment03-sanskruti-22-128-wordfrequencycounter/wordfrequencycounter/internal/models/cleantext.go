package models

import (
	"bytes"
	"unicode"
)


func CleanText(s string) string {
	return string(bytes.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, []byte(s)))
}
