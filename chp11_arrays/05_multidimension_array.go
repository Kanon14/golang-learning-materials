package main

import "fmt"

func main() {

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("Matrix:", matrix[0])

	// If you want to replace an arr within an array, required to recreate
	matrix[0] = [3]int{7, 8, 9}

	fmt.Println("Matrix after change:", matrix[0])

	for _, nested := range matrix {
		for _, value := range nested {
			fmt.Println(value)
		}
	}

}
