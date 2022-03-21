package v3

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	f.Add("")
	f.Add(" ")
	f.Add("hello world!")

	f.Fuzz(func(t *testing.T, s string) {
		if reversed, err := Reverse(s); err == nil {
			assert.True(t, utf8.ValidString(reversed))
			if original, err := Reverse(reversed); err == nil {
				assert.True(t, original == s)
				assert.True(t, utf8.ValidString(original))
			} else {
				assert.ErrorIs(t, err, ErrInvalidString)
			}
		} else {
			assert.ErrorIs(t, err, ErrInvalidString)
		}
	})
}
