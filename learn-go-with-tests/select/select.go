package main

import (
	"fmt"
	"net/http"
	"time"
)

func RacerOld(url1, url2 string) (winner string) {
	tot1 := timeResp(url1)
	tot2 := timeResp(url2)

	if tot1 < tot2 {
		return url1
	}
	return url2

}

func timeResp(url string) time.Duration {
	st := time.Now()
	http.Get(url)
	return time.Since(st)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func Racer(url1, url2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}

}
