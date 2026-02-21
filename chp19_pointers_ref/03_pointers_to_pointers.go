package main

func test(pointerSlice *[]*int) {
	values := *pointerSlice
	for _, value := range values {
		value
	}
}

func main() {
	// 	x := 0
	// 	y := &x
	// 	z := &y
	// 	fmt.Println(x, *y, **z)
	a := 1
	b := 2
	values := &[]*int(&a, &b)
}
