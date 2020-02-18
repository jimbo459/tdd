package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

type Sleeper interface {
	Sleep()
}

const finalWord = "Go!"
const CountStart = 3

func Countdown(w io.Writer, s Sleeper) {
	for i := CountStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}

	s.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{2 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

