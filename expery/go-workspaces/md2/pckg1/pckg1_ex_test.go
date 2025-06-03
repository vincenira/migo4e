package pckg1

import "testing"

func TestCalculateSequence(t *testing.T) {
	got := CalculaterSequenceNumber(10)
	want := 55

	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
