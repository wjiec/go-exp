package v3

import (
	"errors"
	"unicode/utf8"
)

//go:generate go test -v -fuzz=Fuzz -fuzztime=15s

var (
	// ErrInvalidString indicates the passed string is not in the valid utf8 format
	ErrInvalidString = errors.New("invalid string")
)

// Reverse returns a valid utf8 string in reversed, returns error
// when passed string is not valid utf8 format
func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return "", ErrInvalidString
	}

	b := []rune(s)
	for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
	return string(b), nil
}
