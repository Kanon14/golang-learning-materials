package main

import "fmt"

func main() {
	type User struct {
		Name string
		Age  int
	}

	users := map[int]User{
		1: {"Alice", 20},
		2: {"Bob", 25},
	}
	fmt.Printf("%T", users)         // map[int]main.User
	fmt.Printf("%T", users[1].Name) // Alice (string)
}
