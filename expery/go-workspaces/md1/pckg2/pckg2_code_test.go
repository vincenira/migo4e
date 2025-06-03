package pckg2

import "testing"

func TestGreeter(t *testing.T) {
	got := Greeter("Toto")
	want := "Salama Toto"

	if got != want {
		t.Errorf("got \"%v\" but want \"%v\"", got, want)
	}
}
