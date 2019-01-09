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
		got := Hello("Sarah", "English")
		want := "Hello Sarah"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("say 'Hell world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("say greeting in Japanese", func(t *testing.T) {
		got := Hello("Tomoko", "Japanese")
		want := "konnichiwa Tomoko"
		assertCorrectGreeting(t, got, want)
	})

	t.Run("say greeting in German", func(t *testing.T) {
		got := Hello("Barbara", "German")
		want := "Hallo Barbara"
		assertCorrectGreeting(t, got, want)
	})
}
