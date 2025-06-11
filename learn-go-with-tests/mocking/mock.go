package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3
const sleep = "sleep"
const write = "write"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++

}

type ConfSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *ConfSleeper) Sleep() {
	s.sleep(s.duration)

}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type SpyOps struct {
	Calls []string
}

func (s *SpyOps) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyOps) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(out io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()

	}
	fmt.Fprint(out, finalWord)
}

func main() {
	s := &ConfSleeper{5 * time.Second, time.Sleep}
	Countdown(os.Stdout, s)

}
