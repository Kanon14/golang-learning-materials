package main

import "fmt"

func getFunc(str string) func(string) string {
	return func(str2 string) string {
		return str + str2
	}
}

func main() {
	f1 := getFunc("Hello")
	value := f1(" World")
	value2 := f1(" Kanon")
	fmt.Println(value, value2)
}
