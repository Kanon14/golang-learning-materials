package main

import "fmt"

func main() {
	arr := [3]int{100, 200, 300}

	// Classif for-loop
	for i := 0; i < len(arr); i++ {
		fmt.Println("Index:", i, "Value:", arr[i])
	}

	// Recommended: Using range
	for _, value := range arr {
		fmt.Println("Value:", value)
	}

}
