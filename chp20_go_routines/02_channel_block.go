package main

import "fmt"

func add(x int, y int, ch chan int) {
	ch <- x + y // Passing a value to a channel
}

func main() {
	ch := make(chan int) // Make a channel to pass a go process
	// value := go add(5, 10) This does not work
	go add(5, 10, ch)
	x := <-ch // Receive the value from the channel (aka blocking code)
	fmt.Println(x)
}
