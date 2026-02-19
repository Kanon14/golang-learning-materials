package main

import "fmt"

func main() {

	s := []string{"apple", "banana", "cherry"}

	for i, v := range s {
		fmt.Println(i, v)
	}
}
