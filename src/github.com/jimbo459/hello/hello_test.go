package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("James")
	want := "Hello James"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
