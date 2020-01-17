package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying Hello to people", func(t *testing.T) {
		got := hello("James")
		want := "Hello, James"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when empty string provided", func(t *testing.T) {
		got := hello("")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}
