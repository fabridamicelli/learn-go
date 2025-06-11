package main

import (
	"sync"
	"testing"
)

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

func TestHelloWorld(t *testing.T) {
	t.Run("inc 3 times, leaves at 3", func(t *testing.T) {
		c := Counter{}
		c.Inc()
		c.Inc()
		c.Inc()

		assertCounter(t, &c, 3)
	})

	t.Run("safe concurrency", func(t *testing.T) {
		want := 1000
		c := Counter{}

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func() {
				c.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, &c, want)

	})
}
