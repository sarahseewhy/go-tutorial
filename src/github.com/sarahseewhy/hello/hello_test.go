package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectGreeting := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sarah")
		want := "Hello Sarah"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("say 'Hell world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello world"
		assertCorrectGreeting(t, got, want)
	})
}
