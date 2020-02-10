package main

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

const finalWord = "Go!"
const CountStart = 3

func Countdown(w io.Writer) {
	for i := CountStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(w, i)
	}
	fmt.Fprint(w, finalWord)
}

