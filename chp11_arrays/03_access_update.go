package main

import "fmt"

func main() {

	arr := [4]int{10, 20, 30, 40}

	fmt.Println("First element:", arr[0])
	fmt.Println("Third element:", arr[2])

	// Update value
	arr[1] = 99

	fmt.Println("Updated array:", arr)
}
