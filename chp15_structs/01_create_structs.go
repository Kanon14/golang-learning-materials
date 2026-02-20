package main

import "fmt"

type Person struct {
	Name string
	Age  int
	F    func(string) string
}

// func getName(p Person) string {
// 	return p.name
// }

func (p Person) GetName() string {
	return p.Name
}

func (p Person) SetName(newName string) {
	p.Name = newName
	fmt.Println(p)
}

func main() {
	p1 := Person{Name: "Alice", Age: 20}
	p1.SetName("Joey")
	fmt.Println(p1)
}

// Capital letter in struct or function
// allow us to use them outside the package
// Exported name vs unexported name
