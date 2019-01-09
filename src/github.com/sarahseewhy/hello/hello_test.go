package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Sarah")
		want := "Hello Sarah"

		if got != want {
			t.Errorf("got '%s' wnat '%s'", got, want)
		}
	})

	t.Run("say 'Hell world' when an empty string is supplied", func(t *testing.T){

		got := Hello("")
		want := "Hello world"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
