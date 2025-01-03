package series

import (
	"reflect"
	"testing"
)

func checkSubstr(t testing.TB, got, exp []string) {
	t.Helper()
	if !reflect.DeepEqual(got, exp) {
		t.Errorf("Got %q, exp %q", got, exp)
	}
}

func TestAll(t *testing.T) {
	cases := []struct {
		name string
		n    int
		s    string
		exp  []string
	}{
		{"simple n=3", 3, "49142", []string{"491", "914", "142"}},
		{"n>len(s)", 8, "49142", []string{"49142"}},
		{"n=len(s)", 5, "49142", []string{"49142"}},
		{"empty string", 5, "", []string{""}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := All(c.n, c.s)
			exp := c.exp
			checkSubstr(t, got, exp)
		})
	}
}
