package main

import "fmt"

func main() {
	a := 1

	switch {
	case a < -1:
		fmt.Println("a is less than -1")
	case a < 0:
		fmt.Println("a is less than 0")
	case a < 1:
		fmt.Println("a is less than 1")
	default:
		fmt.Println("default")
	}
}
