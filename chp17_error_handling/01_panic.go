package main

import "fmt"

func divide(a int, b int) int {
	return a / b
}

func deferFunc() {
	r := recover()
	if r != nil {
		fmt.Println("Error:", r)
	} else {
		fmt.Println("No error")
	}
}

func main() {
	defer deferFunc() // This allow the code to run after the main function finishes
	panic("this is a crash")
	// panic is a runtime error
	// anything below the occurred panic will not be executed.

	fmt.Println("run")
}
