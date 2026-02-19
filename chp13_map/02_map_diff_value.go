package main

import "fmt"

func main() {

	// Pattern map[key][]dtype
	mp := map[string][]int{"a": {1, 2, 3}}
	mp["b"] = []int{4, 5, 6}
	mp["c"] = []int{7, 8, 9}
	delete(mp, "b")
	fmt.Println(mp)
}
