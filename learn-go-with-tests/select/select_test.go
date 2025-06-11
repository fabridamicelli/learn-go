package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {

	t.Run("compare Servers speed", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL, 1*time.Second)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("check timeout", func(t *testing.T) {
		serv1 := makeDelayedServer(2 * time.Second)
		serv2 := makeDelayedServer(3 * time.Second)
		defer serv1.Close()
		defer serv2.Close()
		_, err := Racer(serv1.URL, serv2.URL, 1*time.Second)

		if err == nil {
			t.Error("Expected error, no errors found")
		}
	})

}
