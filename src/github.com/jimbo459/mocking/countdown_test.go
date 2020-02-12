package main

import (
	"testing"
	"reflect"
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

	if !reflect.DeepEqual(want, spySleepPrinter.Calls){
		t.Errorf("wanted calls %v, got %v", want, spySleepPrinter.Calls)
	}

}
