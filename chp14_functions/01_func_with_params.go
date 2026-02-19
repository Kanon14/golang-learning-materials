package main

import "fmt"

func add(a int, b int) (int, string) {
	return a + b, "this is function operation in Go"
}

func main() {
	sum, str := add(10, 20)
	fmt.Println(sum, str)

}
