package main

import "fmt"

type Sport struct {
	name     string
	position string
}

type Person struct {
	name     string
	age      uint
	favSport []Sport // Embed struct into Person struct
}

func main() {
	p1 := Person{age: 23, name: "Kanon", favSport: []Sport{"Soccer", "RW"}}
	fmt.Println(p1.favSport.position)
}
