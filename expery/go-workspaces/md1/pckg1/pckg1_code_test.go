package pckg1

import (
	"testing"
)

func TestAdder(t *testing.T) {
	got := Adder(4, 5)
	assertHelper(t, got, 9)
}

func TestSubtracter(t *testing.T) {
	got := Subtracter(10, 4)
	assertHelper(t, got, 6)
}

func assertHelper(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("Got %d but Want %d", got, want)
	}
}
