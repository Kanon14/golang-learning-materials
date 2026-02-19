package main

import "fmt"

func sum(nums ...int) (total int, multi int) {

	for _, n := range nums {
		total += n
		multi += 2 * n
	}
	return total, multi
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(sum([]int{1, 2, 3, 4, 5}...))
}
