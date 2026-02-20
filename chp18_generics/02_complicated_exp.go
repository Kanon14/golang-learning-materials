package main

import "fmt"

func getValues[K comparable, V any](mp map[K]V) []V {
	values := []V{}

	for _, value := range mp {
		values = append(values, value)
	}
	return values
}

func main() {
	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	values := getValues(mp)
	fmt.Println(values)
}
