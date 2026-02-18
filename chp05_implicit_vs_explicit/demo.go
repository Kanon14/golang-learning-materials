package main

import "fmt"

// This is an explicit method
// func main() {
// 	y := 3.5            // Implicitly assign the method, auto data type
// 	fmt.Printf("%T", y) // Print data type
// }

// func main() {
// 	y := uint(0) // Type casting
// 	fmt.Printf("%T", y)
// }

func main() {
	x := -9
	y := uint(x) // This will be an error, since we convert int to uint
	fmt.Println(x, y)
}
