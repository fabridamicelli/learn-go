package runlength

import (
	"testing"
)

func TestRunningLengthEncoding(t *testing.T) {
	cases := []struct {
		name string
		s    string
		exp  string
	}{
		{"empty", "", ""},
		{"1char", "a", "a"},
		{"aaa", "aaa", "3a"},
		{"aab", "aab", "2ab"},
		{"abc", "abc", "abc"},
		{"aabbcc", "aabbcc", "2a2b2c"},
		{"aabbcca", "aabbcca", "2a2b2ca"},
		{"aabcc", "aabcc", "2ab2c"},
		{"long", "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB", "12WB12W3B24WB"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := RunningLengthEncode(c.s)
			if got != c.exp {
				t.Fatalf("Got %s, expected %s", got, c.exp)
			}

		})
	}

}

func TestRunningLengthDecoding(t *testing.T) {
	cases := []struct {
		name string
		s    string
		exp  string
	}{
		{"empty", "", ""},
		{"3a", "3a", "aaa"},
		{"2a1b", "2a1b", "aab"},
		{"1a1b1c", "1a1b1c", "abc"},
		{"2a2b2c", "2a2b2c", "aabbcc"},
		{"2a2b2c1a", "2a2b2c1a", "aabbcca"},
		{"10a", "10a", "aaaaaaaaaa"},
		{"2AB3CD4E", "2AB3CD4E", "AABCCCDEEEE"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, _ := RunningLengthDecode(c.s)
			if got != c.exp {
				t.Fatalf("Got %s, expected %s", got, c.exp)
			}

		})
	}

}
