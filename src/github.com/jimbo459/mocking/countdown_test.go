package main

import (
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type CountdownOperationsSpy struct {
	Calls []string
}

func (cs *CountdownOperationsSpy) Sleep() {
	cs.Calls = append(cs.Calls, sleep)
}

func (cs *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	cs.Calls = append(cs.Calls, write)
	return
}

type SpyTime struct {
	timeSlept time.Duration
}

func (st *SpyTime) Sleep(duration time.Duration) {
	st.timeSlept = duration
}

func TestCountdown(t *testing.T) {

	spySleepPrinter := &CountdownOperationsSpy{}
	Countdown(spySleepPrinter, spySleepPrinter)

	want := []string{
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
		t.Errorf("wanted calls %v, got %v", want, spySleepPrinter.Calls)
	}

}

func TestConfigurableSleeper(t *testing.T) {
	expectedDuration := 5 * time.Second

	timeSpy := &SpyTime{}

	sleeper := &ConfigurableSleeper{expectedDuration, timeSpy.Sleep}

	sleeper.Sleep()

	if expectedDuration != timeSpy.timeSlept {
		t.Errorf("expected %v, got %v", expectedDuration, timeSpy.timeSlept)
	}

}
