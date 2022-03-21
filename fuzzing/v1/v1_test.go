/*
Fuzz tests must be in *_test.go files to run.
*/

package v1

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

// A fuzz test must be a function named like FuzzXxx, which accepts only a *testing.F, and has no return value.
func FuzzReverse(f *testing.F) {
	// All seed corpus entries must have types which are identical to the fuzzing arguments, in the same order.
	//This is true for calls to (*testing.F).Add and any corpus files in the testdata/fuzz directory of the fuzz test.
	f.Add("")
	f.Add(" ")
	f.Add("hello world!")

	// There must be exactly one fuzz target per fuzz test.
	//
	// A fuzz target must be a method call to (*testing.F).Fuzz which accepts a *testing.T as the first parameter,
	// followed by the fuzzing arguments. There is no return value.
	f.Fuzz(func(t *testing.T, s string) {
		reversed := Reverse(s)
		assert.True(t, utf8.ValidString(reversed))
		assert.True(t, Reverse(reversed) == s)
	})
}
