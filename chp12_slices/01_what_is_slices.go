package main

import "fmt"

func main() {

	// Slice literal (no fixed size)
	numbers := []int{10, 20, 30}

	fmt.Println(numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
}
