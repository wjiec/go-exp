package v1

//go:generate go test -v -fuzz=Fuzz

// Reverse returns a string in reversed
func Reverse(s string) string {
	b := []byte(s)
	for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
	return string(b)
}

/*
failed corpus file:

go test fuzz v1
string("\xff")
*/
