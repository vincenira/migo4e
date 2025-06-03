package pckg1

import "github.com/vincenira/migo4e/expery/go-workspaces/md1/pckg1"

func CalculaterSequenceNumber(n int) int {
	return (n * pckg1.Adder(n, 1)) / 2
}
