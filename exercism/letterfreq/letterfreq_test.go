package letterfreq

import (
	"reflect"
	"testing"
)

func TestLetterFreq(t *testing.T) {
	cases := []struct {
		name  string
		texts []string
		exp   []map[rune]int
	}{
		{"empty", []string{"", ""}, []map[rune]int{{}, {}}},
		{"abb,cdd", []string{"ab b", "c dd"}, []map[rune]int{{'a': 1, 'b': 2}, {'c': 1, 'd': 2}}},
		{"abb,cdd", []string{"ab\nb", "c\tdd"}, []map[rune]int{{'a': 1, 'b': 2}, {'c': 1, 'd': 2}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := LetterFreq(c.texts)
			if !reflect.DeepEqual(got, c.exp) {
				t.Fatalf("Got %v, exp %v", got, c.exp)
			}
		})
	}

}
