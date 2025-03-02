package main

import "testing"

func TestHelloK(t *testing.T) {
	want := "kashikuro"
	if ggot := helloK(); ggot != want {
		t.Errorf(
			"helloK() = %v, want %v", ggot, want)
	}
}
