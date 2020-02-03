package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "James")

	got := buffer.String()
	want := "Hello, James"

	if got != want {
		t.Errorf("Want %s, got %s", want, got)
	}
}
