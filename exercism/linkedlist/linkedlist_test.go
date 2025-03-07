package linkedlist_test

import (
	"exercism/linkedlist"
	"reflect"
	"testing"
)

func TestList(t *testing.T) {

	t.Run("one element", func(t *testing.T) {
		l1 := linkedlist.NewList("hello")
		got := l1.First().Value
		exp := "hello"
		if got != exp {
			t.Errorf("got %s, expected %s", got, exp)
		}

		if l1.First().Value != "hello" {
			t.Errorf("First value should be hello, got %s", l1.First().Value)
		}

		if l1.Last().Value != "hello" {
			t.Errorf("Last value should be hello, got %s", l1.First().Value)
		}

		if l1.First().Prev() != nil {
			t.Errorf("Prev should return nil")
		}

		if l1.First().Next() != nil {
			t.Errorf("Next should return nil")
		}

	})

	l := linkedlist.NewList(1, 2, 3, 4)

	assertVal := func(t testing.TB, got, exp interface{}) {
		t.Helper()
		if got != exp {
			t.Errorf("got %d, expected %d", got, exp)
		}

	}

	t.Run("list first method", func(t *testing.T) {
		got := l.First().Value
		exp := 1
		assertVal(t, got, exp)
	})

	t.Run("list last method", func(t *testing.T) {
		got := l.Last().Value
		exp := 4
		assertVal(t, got, exp)
	})

	t.Run("node next method", func(t *testing.T) {
		n := l.First()
		for exp := 2; exp < 5; exp++ {
			n = n.Next()
			got := n.Value
			assertVal(t, got, exp)

			if exp == 4 {
				last := n.Next()
				if last != nil {
					t.Fatalf("Expected %v, got %v", nil, last)

				}
			}
		}
	})

	t.Run("node prev method", func(t *testing.T) {
		n := l.Last()
		for exp := 3; exp > 0; exp-- {
			n = n.Prev()
			got := n.Value
			assertVal(t, got, exp)

			if exp == 1 {
				first := n.Prev()
				if first != nil {
					t.Fatalf("Expected %v, got %v", nil, first)
				}
			}
		}
	})

	type testStruct struct {
		a string
		b int
	}

	t.Run("list Push method", func(t *testing.T) {
		ll := linkedlist.NewList(1, 2, 3)
		ll.Push(4)
		assertVal(t, ll.First().Value, 4)
		ll.Push(22)
		assertVal(t, ll.First().Value, 22)

		ll.Push(testStruct{
			a: "hello",
			b: 121,
		},
		)
		exp := testStruct{a: "hello", b: 121}
		got := ll.First().Value
		if !reflect.DeepEqual(got, exp) {
			t.Fatalf("Expected %v, got %v", exp, got)
		}

	})

	t.Run("list Unshift method", func(t *testing.T) {
		ll := linkedlist.NewList(1, 2, 3)
		ll.Unshift(4)
		assertVal(t, ll.Last().Value, 4)
		ll.Unshift(22)
		assertVal(t, ll.Last().Value, 22)

		type testStruct struct {
			a string
			b int
		}

		ll.Unshift(testStruct{
			a: "hello",
			b: 121,
		},
		)
		exp := testStruct{a: "hello", b: 121}
		got := ll.Last().Value
		if !reflect.DeepEqual(got, exp) {
			t.Fatalf("Expected %v, got %v", exp, got)
		}

	})

	t.Run("list Pop method", func(t *testing.T) {
		ll := linkedlist.NewList(1, 2, 3, testStruct{a: "ciao", b: 1})
		out, _ := ll.Pop()
		assertVal(t, out, 1)
		assertVal(t, ll.First().Value, 2)

		out, _ = ll.Pop()
		assertVal(t, out, 2)
		assertVal(t, ll.First().Value, 3)

		exp := testStruct{a: "ciao", b: 1}
		got := ll.Last().Value
		if !reflect.DeepEqual(got, exp) {
			t.Fatalf("Expected %v, got %v", exp, got)
		}
		ll.Pop()
		ll.Pop() // make list empty

		if ll.NumNodes != 0 {
			t.Fatalf("Expected 0, got %d", ll.NumNodes)
		}
		if ll.First() != nil {
			t.Fatalf("Expected nil, got %v", ll.First())
		}
		if ll.Last() != nil {
			t.Fatalf("Expected nil, got %v", ll.Last())
		}

	})

	t.Run("list Shift method", func(t *testing.T) {
		ll := linkedlist.NewList(1, 2, 3, testStruct{a: "ciao", b: 1})
		got, _ := ll.Shift()
		exp := testStruct{a: "ciao", b: 1}
		if !reflect.DeepEqual(got, exp) {
			t.Fatalf("Expected %v, got %v", exp, got)
		}
		assertVal(t, ll.Last().Value, 3)
		ll.Shift()
		assertVal(t, ll.Last().Value, 2)
		ll.Shift()
		assertVal(t, ll.Last().Value, 1)
		ll.Shift()
		v, _ := ll.Shift() // empty the list
		if v != nil {
			t.Fatalf("Expected nil, got %v", v)
		}

	})

	t.Run("list Reverse method", func(t *testing.T) {

		l1 := linkedlist.NewList(1)
		l1.Reverse()
		assertVal(t, l1.First().Value, 1)

		ll := linkedlist.NewList(1, 2, 3, 4)
		ll.Reverse()
		assertVal(t, ll.Last().Value, 1)
		assertVal(t, ll.First().Value, 4)

		assertVal(t, ll.First().Next().Value, 3)
		assertVal(t, ll.First().Next().Next().Value, 2)
		assertVal(t, ll.First().Next().Next().Next().Value, 1)

	})
}
