package hamming

import (
	"fmt"
	"testing"
)

func TestHamming(t *testing.T) {
	cases := []struct {
		name string
		a    string
		b    string
		exp  int
	}{
		{"empty", "", "", 0},
		{"nodiff", "aa", "aa", 0},
		{"1diff", "aa", "ab", 1},
		{"alldiff", "aa", "bb", 2},
		{"alldiff", "abababab", "acacacac", 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, _ := Hamming(c.a, c.b)
			if got != c.exp {
				t.Errorf("Inp: (%s,%s), Got %d, expected %d", c.a, c.b, got, c.exp)
			}
		})
	}

}
func TestHammingInvalid(t *testing.T) {
	cases := []struct {
		name string
		a    string
		b    string
		exp  error
	}{
		{"!= len", "2", "", fmt.Errorf("Undefined")},
		{"!= len", "22", "2", fmt.Errorf("Undefined")},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := Hamming(c.a, c.b)
			if err == nil {
				t.Errorf("Expected error")
			}
		})
	}

}
