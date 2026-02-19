package main

import "fmt"

// func main() {

// 	m := make(map[string]int)

// 	m["apple"] = 10
// 	m["banana"] = 20

// 	fmt.Println(m)
// }

func main() {

	m := map[string]int{
		"apple":  10,
		"banana": 20,
	}

	fmt.Println(m)

	// Accessing Values
	fmt.Println(m["apple"])  // 10
	fmt.Println(m["orange"]) // 0 (zero value!)

	// Replace
	m["banana"] = 15
	fmt.Println(m)
}
