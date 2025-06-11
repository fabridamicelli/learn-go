package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("with a", func(t *testing.T) {
		got := Repeat("a", 5)
		exp := "aaaaa"
		if got != exp {
			t.Errorf("Got %s, expected %s", got, exp)
		}

	})
	t.Run("with p", func(t *testing.T) {
		got := Repeat("p", 3)
		exp := "ppp"
		if got != exp {
			t.Errorf("Got %s, expected %s", got, exp)
		}

	})

}

func TestRepeatStdLib(t *testing.T) {
	t.Run("with a", func(t *testing.T) {
		got := RepeatStdLib("a", 5)
		exp := "aaaaa"
		if got != exp {
			t.Errorf("Got %s, expected %s", got, exp)
		}

	})
	t.Run("with p", func(t *testing.T) {
		got := RepeatStdLib("p", 3)
		exp := "ppp"
		if got != exp {
			t.Errorf("Got %s, expected %s", got, exp)
		}

	})

}

func ExampleRepeat() {
	out := Repeat("f", 4)
	fmt.Println(out)
	// Output: ffff
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}

}

func BenchmarkRepeatSdtLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatStdLib("a", 10)
	}

}
