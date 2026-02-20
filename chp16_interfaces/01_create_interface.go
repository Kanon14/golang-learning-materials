package main

import "fmt"

type Shape interface {
	getPerimeter() uint
}

type Square struct {
	width uint
}

func (s Square) getPerimeter() uint {
	return 4 * s.width
}

type Triangle struct {
	a uint
	b uint
	c uint
}

func (t Triangle) getPerimeter() uint {
	return t.a + t.b + t.c
}

func (t Triangle) getSides() []uint {
	return []uint{t.a, t.b, t.c}
}

// Note: Cannot directly instantiate an interface

func main() {
	var shapes []Shape = []Shape{Triangle{1, 2, 3}, Square{10}}
	// fmt.Println(s.getPerimeter()) // Note that we only able to access the getPerimeter()
	// fmt.Println(s.getSides()) // This will cause an error

	perimeters := uint(0)

	for _, shpae := range shapes {
		perimeters += shpae.getPerimeter()
	}

	fmt.Println(perimeters)
}
