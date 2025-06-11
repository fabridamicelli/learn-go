package main

import "testing"

func assertCorrectMessage(t testing.TB, got, exp string) {
	t.Helper()
	if got != exp {
		t.Errorf("got %q, want %q", got, exp)
	}

}

func TestCheckLang(t *testing.T) {
	t.Run("valid lang", func(t *testing.T) {
		_, got := checkLang("es")
		exp := "es"
		assertCorrectMessage(t, got, exp)
	})

	t.Run("invalid lang", func(t *testing.T) {
		err, _ := checkLang("espanol")
		if err == nil {
			t.Errorf("expected err, got nil")
		}
	})

}

func TestHello(t *testing.T) {
	t.Run("with name", func(t *testing.T) {
		got := Hello("Fab", "")
		exp := "Hello, Fab"
		assertCorrectMessage(t, got, exp)

	})

	t.Run("without name", func(t *testing.T) {
		got := Hello("", "")
		exp := "Hello, world"
		assertCorrectMessage(t, got, exp)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "es")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in English", func(t *testing.T) {
		got := Hello("Elodie", "en")
		want := "Hello, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in German", func(t *testing.T) {
		got := Hello("Elodie", "de")
		want := "Hallo, Elodie"
		assertCorrectMessage(t, got, want)
	})

}
