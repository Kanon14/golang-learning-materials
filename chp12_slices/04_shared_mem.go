package main

import "fmt"

func main() {

	arr := [4]int{1, 2, 3, 4}

	s := arr[1:3]
	s[0] = 99

	fmt.Println("Array:", arr)
	fmt.Println("Slice:", s)
}
