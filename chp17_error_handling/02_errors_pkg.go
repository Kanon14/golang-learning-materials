package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by 0")
	}
	return a / b, nil
}

func main() {
	result, err := divide(1, 0)
	fmt.Println(result, err)
}
