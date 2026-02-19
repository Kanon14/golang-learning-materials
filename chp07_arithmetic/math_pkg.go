package main

import (
	"fmt"
	"math"
)

func main() {

	x := 16.0
	y := 2.5

	fmt.Println("Square Root:", math.Sqrt(x))
	fmt.Println("Power:", math.Pow(2, 3))
	fmt.Println("Ceil:", math.Ceil(y))
	fmt.Println("Floor:", math.Floor(y))
	fmt.Println("Round:", math.Round(y))
	fmt.Println("Absolute:", math.Abs(-10))
	fmt.Println("Max:", math.Max(5, 9))
	fmt.Println("Min:", math.Min(5, 9))
}
