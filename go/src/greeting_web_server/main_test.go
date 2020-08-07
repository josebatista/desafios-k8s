package main

import "testing"

func TestGreeting(t *testing.T) {

	expected := "<b>Code.education Rocks!</b>"

	received := greeting("Code.education Rocks!")

	if received != expected {
		t.Errorf("Resposta incorreta. Esperava %s e recebeu %s", expected, received)
	}
}
