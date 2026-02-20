package main

import "fmt"

type Number interface {
	int | float64 | uint
}

// T is the union of int and float64, only both are int / float will work
func add[T Number](x T, y T) T {
	var sum T = x + y
	return sum
}

func main() {
	v1 := add(1, 2)
	v2 := add(1.0, 2.0)
	fmt.Println(v1, v2)
}
