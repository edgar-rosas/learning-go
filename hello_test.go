package main

import "testing"

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("greeting people", func(t *testing.T) {
		got := Hello("Edgar")
		want := "Hello, Edgar!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greet the world when empty string is found", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})
}
