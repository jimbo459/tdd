package main

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

const finalWord = "Go!"
const CountStart = 3

func Countdown(w io.Writer, s Sleeper) {
	for i := CountStart; i > 0; i-- {
		s.Sleep()
	}
	for i := CountStart; i > 0; i-- {
		fmt.Fprintln(w, i)
	}

	s.Sleep()
	fmt.Fprint(w, finalWord)
}

