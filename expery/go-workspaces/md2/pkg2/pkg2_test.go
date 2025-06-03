package pkg2

import "testing"

func TestAdderMrMrs(t *testing.T) {
	got := AdderMrMrs("toto", 'm')
	want := "Salama Mr toto"

	if got != want {
		t.Errorf("got \"%v\" but want \"%v\"", got, want)
	}
	cases := []struct {
		testName string
		name     string
		char     byte
		want     string
	}{
		{"test1", "toto", 'm', "Salama Mr toto"},
		{"test2", "titi", 'f', "Salama Mrs titi"},
		{"test3", "tutu", 'n', "Salama Mr/Mrs tutu"},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			got := AdderMrMrs(c.name, c.char)
			AssertHelperAdder(t, got, c.want)
		})
	}
}

func AssertHelperAdder(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got \"%v\" but want \"%v\"", got, want)
	}
}
