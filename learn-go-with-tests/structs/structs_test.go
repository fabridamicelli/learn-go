package main

import "testing"

func checkNumber(t testing.TB, got, exp float64) {
	t.Helper()
	if got != exp {
		t.Errorf("Got %g, exp %g", got, exp)
	}
}

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 10.0}
	got := Perimeter(rec)
	exp := 40.0
	checkNumber(t, got, exp)

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		exp   float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		exp := tt.exp
		checkNumber(t, got, exp)
	}
}
