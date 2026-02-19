package main

import "fmt"

func main() {

	num := 5

	num += 3         // num = num + 3
	fmt.Println(num) // 8

	num -= 2
	fmt.Println(num) // 6

	num *= 4
	fmt.Println(num) // 24

	num /= 3
	fmt.Println(num) // 8

	num %= 5
	fmt.Println(num) // 3
}
