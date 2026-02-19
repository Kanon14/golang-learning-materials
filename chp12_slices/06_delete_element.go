package main

import "fmt"

func main() {

	s := []int{10, 20, 30, 40, 50}

	index := 2 // remove value 30

	s = append(s[:index], s[index+1:]...)

	fmt.Println(s)
}
