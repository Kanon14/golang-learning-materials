package main

import "fmt"

func main() {

	var s1 []int  // nil slice
	s2 := []int{} // empty slice

	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil) // false

	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
}
