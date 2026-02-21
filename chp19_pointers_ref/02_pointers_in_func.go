package main

import "fmt"

// func change(x *int) {
// 	*x = 100
// }

type Book struct {
	id    int
	title string
} // Derefence only work on struct

func (b *Book) setTitle(title string) {
	b.title = title
}

func main() {
	// a := 10
	// change(&a)
	// fmt.Println(a)

	b := Book{id: 1, title: "Go Programming"}
	fmt.Println(b)
	b.setTitle("New to Go Programming")
	fmt.Println(b)
}
