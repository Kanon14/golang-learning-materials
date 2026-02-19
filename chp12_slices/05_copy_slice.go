package main

import "fmt"

func main() {

	original := []int{10, 20, 30}
	copySlice := make([]int, len(original))

	copy(copySlice, original)

	// Only copy being altered
	copySlice[0] = 99

	fmt.Println("Original:", original)
	fmt.Println("Copied:", copySlice)
}
