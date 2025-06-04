package stringutils

import (
	"errors"
	"unicode/utf8"
)

// UnSafeReverseString reverses a string by bytes. This is unsafe for multibyte
// UTF-8 characters and may corrupt them.
func UnSafeReverseString(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// ReverseString safely reverses a string by runes. This preserves multibyte
// UTF-8 characters.
func ReverseString(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
