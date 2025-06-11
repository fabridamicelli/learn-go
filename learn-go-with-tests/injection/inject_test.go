package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet("Fabri", &buffer)

	got := buffer.String()
	exp := "Hello, Fabri"

	if got != exp {
		t.Errorf("expected %s, got %s", exp, got)
	}
}
