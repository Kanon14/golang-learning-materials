package main

import "fmt"

func main() {

	arr := [6]int{10, 20, 30, 40, 50, 60}

	// Slice expression: [low:high)
	s1 := arr[1:4] // index 1 to 3
	fmt.Println(s1)

	// From another slice
	s2 := s1[1:3]
	fmt.Println(s2)
}
