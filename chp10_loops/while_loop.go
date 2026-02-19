package main

import "fmt"

// In Golang, there is no while loop like in js or python

// func main() {
// 	a := 1
// 	for a < 10 {
// 		fmt.Println("loop")
// 		a++
// 	}
// }

// func main() {
// 	str := "hello world"
// 	for idx := 0; idx < 10; idx++ {
// 		fmt.Printf("%T", str[idx])
// 	} // -> uint8 = byte
// }

func main() {
	str := "hello world"
	for _, char := range str {
		fmt.Printf("%T", char)
	} // -> int32 = rune
}
