package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {

	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}

	Countdown(buffer, sleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if sleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", sleeper.Calls)
	}

}
