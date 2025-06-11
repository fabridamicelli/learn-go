package main

import "testing"

func assertStrings(t testing.TB, got, exp string) {
	t.Helper()
	if got != exp {
		t.Errorf("got %s, expected %s", got, exp)
	}

}

func assertError(t testing.TB, got, exp error) {
	t.Helper()
	if got != exp {
		t.Errorf("got %s, expected %s", got, exp)
	}

}

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		d := Dict{"test": "just a test"}
		got, _ := d.Search("test")
		exp := "just a test"
		assertStrings(t, got, exp)
	})

	t.Run("unknown word", func(t *testing.T) {
		d := Dict{"test": "just a test"}
		_, err := d.Search("foo")

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		assertStrings(t, err.Error(), WordNotFoundError.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		d := Dict{"test": "just a test"}
		d.Add("new", "funny definition")
		got, _ := d.Search("new")
		exp := "funny definition"
		assertStrings(t, got, exp)
	})

	t.Run("add existing word", func(t *testing.T) {
		d := Dict{"test": "just a test"}
		err := d.Add("test", "funny definition")
		assertError(t, err, WordAlreadyExistsError)
		got, _ := d.Search("test")
		assertStrings(t, got, "just a test")
	})

	t.Run("update", func(t *testing.T) {
		d := Dict{"test": "just a test"}
		newDef := "funny definition"
		d.Update("test", newDef)
		got, _ := d.Search("test")
		assertStrings(t, got, newDef)
	})

	t.Run("delete", func(t *testing.T) {
		w := "test"
		d := Dict{w: "just a test"}
		d.Delete(w)
		_, err := d.Search(w)
		assertError(t, err, WordNotFoundError)
	})
}
