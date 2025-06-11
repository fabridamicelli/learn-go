package integers

import (
	"fmt"
	"testing"
)

func assertEq(t testing.TB, got, exp int) {
	t.Helper()
	if exp != got {
		t.Errorf("Expected %d, got %d", exp, got)
	}
}

func TestAdder(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		got := Add(2, 2)
		exp := 4
		assertEq(t, got, exp)

	})

	t.Run("negative", func(t *testing.T) {
		got := Add(-2, 2)
		exp := 0
		assertEq(t, got, exp)

	})
}

func ExampleAdd() {
	sum := Add(1, 3)
	fmt.Println(sum)
	// Output: 4

}
