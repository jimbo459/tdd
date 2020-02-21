package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	startA := time.Now()
	http.Get(a)
	aTime := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bTime := time.Since(startB)

	if aTime < bTime {
		return a
	}

	return b
}
