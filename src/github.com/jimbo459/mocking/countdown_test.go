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
	durationSlept time.Duration
}

func (st *SpyTime) Sleep(duration time.Duration) {
	st.durationSlept = duration
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
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if sleepTime != spyTime.durationSlept {
		t.Errorf("wanted %v, got %v", sleepTime, spyTime.durationSlept)
	}

}
