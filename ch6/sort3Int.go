/*
Can you write a function that sorts three int values? try to write two versions of the func-
tion: one with named returned values and another without named return values. Which
one do you think is better?
*/
package main

import "fmt"

func nameReturnTriple(a, b, c int) (s, m, M int) {
	if a < b && a < c {
		s = a
		if b < c {
			M = c
			m = b
		} else {
			M = b
			m = c
		}
	}

	if b < a && b < c {
		s = b
		if a < c {
			M = c
			m = a
		} else {
			M = a
			m = c
		}
	}
	if c < a && c < a {
		s = c
		if b < a {
			M = a
			m = b
		} else {
			M = b
			m = a
		}
	}
	return
}

func unnamed(a, b, c int) (x, y, z int) {
	if a < b && b < c {
		return a, b, c
	}
	if b < a && a < c {
		return b, a, c
	}

	if c < b && b < a {
		return c, b, a
	}
	if b < c && c < a {
		return b, c, a
	}
	return c, a, b
}

func main() {
	x, y, z := unnamed(7, 5, 4)
	a, b, c := nameReturnTriple(7, 5, 4)
	fmt.Println("unnamed: ", x, y, z)
	fmt.Println("named: ", a, b, c)
	x, y, z = unnamed(6, 8, 1)
	a, b, c = nameReturnTriple(6, 8, 1)
	fmt.Println("unnamed: ", x, y, z)
	fmt.Println("named: ", a, b, c)
	x, y, z = unnamed(1, 2, 3)
	a, b, c = nameReturnTriple(1, 2, 3)
	fmt.Println("unnamed: ", x, y, z)
	fmt.Println("named: ", a, b, c)
	x, y, z = unnamed(10, 5, 12)
	a, b, c = nameReturnTriple(10, 5, 12)
	fmt.Println("unnamed: ", x, y, z)
	fmt.Println("named: ", a, b, c)
	x, y, z = unnamed(3, 2, 2)
	a, b, c = nameReturnTriple(3, 2, 2)
	fmt.Println("unnamed: ", x, y, z)
	fmt.Println("named: ", a, b, c)
}
