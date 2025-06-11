package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	t.Run("Prints 3 to Go", func(t *testing.T) {

		buf := &bytes.Buffer{}
		s := &SpySleeper{}

		Countdown(buf, s)

		got := buf.String()
		exp := `3
2
1
Go!`

		if got != exp {
			t.Errorf("got %q, want %q", got, exp)
		}
		if s.Calls != 3 {
			t.Errorf("not enough sleeper calls, want 3, got %d", s.Calls)
		}

	})
	t.Run("check Ops order", func(t *testing.T) {
		spySleepPrinter := &SpyOps{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

}

func TestConfSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept %v, but it slept %v", sleepTime, spyTime.durationSlept)
	}
}
