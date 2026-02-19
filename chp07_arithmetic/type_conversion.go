package main

import (
	"fmt"
	"math"
)

func main() {

	a := 10
	b := 3

	// Convert int to float64
	result := float64(a) / float64(b)

	fmt.Println("Accurate Division:", result)

	// Use math function
	fmt.Println("Square Root of result:", math.Sqrt(result))
}
