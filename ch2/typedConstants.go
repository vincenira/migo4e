// typedConstant has a different data type than i := int(1)
// solution is to cast the i into int16()
// Exercise: correct the error in typedConstants.go

package main

import "fmt"

const (
	typedConstant   = int16(100)
	untypedConstant = 100
)

func main() {
	i := int(1)
	fmt.Println("unTyped:", i*untypedConstant)
	fmt.Println("Typed:", int16(i)*typedConstant)
}
