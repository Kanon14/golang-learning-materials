package main

import (
	"fmt"
)

func main() {

	a := 10
	b := 3

	// Basic arithmetic
	fmt.Println("Addition:", a+b)       // 13
	fmt.Println("Subtraction:", a-b)    // 7
	fmt.Println("Multiplication:", a*b) // 30
	fmt.Println("Division:", a/b)       // 3 (integer division!)
	fmt.Println("Modulus:", a%b)        // 1 (remainder)

	// Float division
	x := 10.0
	y := 3.0
	fmt.Println("Float Division:", x/y) // 3.333...
}
