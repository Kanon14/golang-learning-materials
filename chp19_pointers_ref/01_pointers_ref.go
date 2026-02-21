package main

import "fmt"

func main() {
	x := 0
	y := &x // y is the address of x (hexadecimals)

	*y = 100 // Go to the location of x

	fmt.Println(x, *y) // * asterisk used to dereference
}
