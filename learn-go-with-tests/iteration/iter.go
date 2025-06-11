package iteration

import "strings"

// Return string s repeated n times
func Repeat(s string, n int) string {
	var out string
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}

func RepeatStdLib(s string, n int) string {
	return strings.Repeat(s, n)
}
