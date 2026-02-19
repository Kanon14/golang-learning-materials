package main

import "fmt"

func callFunc(callable func(int) int) int {
	return callable(10)
}

func doubleNumber(num int) int {
	return num * 2
}

func main() {
	// value := callFunc(doubleNumber)
	// fmt.Println(value)

	// // Anonymous Function (Lambda in Python)
	// result := func(a int, b int) int {
	// 	return a * b
	// }(3, 4)
	// fmt.Println("Anoymous Attempt", result)

	f1 := func(x int) int {
		return x + 10
	}
	value := callFunc(f1)
	fmt.Println("Calling Anoymous Function", value)
}
