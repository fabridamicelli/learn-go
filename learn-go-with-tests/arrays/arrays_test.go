package main

import (
	"slices"
	"testing"
)

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("array 5", func(t *testing.T) {
		nums := [5]int{1, 2, 3, 4, 5}
		got := SumArray5(nums)
		exp := 15

		if got != exp {
			t.Errorf("got %d, expected %d", got, exp)

		}

	})
	t.Run("slice", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		got := Sum(nums)
		exp := 10

		if got != exp {
			t.Errorf("got %d, expected %d", got, exp)

		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		exp := []int{3, 9}

		if !slices.Equal(got, exp) {
			t.Errorf("got %d, expected %d", got, exp)

		}
	})

}
func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, exp []int) {
		t.Helper()
		if !slices.Equal(got, exp) {
			t.Errorf("got %d, expected %d", got, exp)
		}

	}

	t.Run("tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		exp := []int{2, 9}
		checkSums(t, got, exp)
	})

	t.Run("tails with empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		exp := []int{0, 9}
		checkSums(t, got, exp)
	})
}

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
	AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
}
