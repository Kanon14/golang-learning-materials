package main

import "fmt"

func main() {
	// Method 1: Declare & Assign Later
	// var nums [3]int
	// nums[0] = 10
	// nums[1] = 20
	// nums[2] = 30

	// fmt.Println(nums)

	// Method 2: Initialize Directly
	// arr := [3]int{1, 2, 3}
	// fmt.Println(arr)

	// Method 3: Let Golang Auto-Count
	auto := [...]int{5, 10, 15, 20, 25}
	fmt.Println(auto)
}
