package main

import (
	"fmt"
	"time"
)

func add(x int, y int, ch chan int, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println(x + y)
	ch <- x + y
}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	// go add(5, 10, ch)
	// go add(8, 11, ch)
	// go add(5, 0, ch)
	// go add(5, -5, ch)
	// x := <-ch
	// fmt.Println(x)

	go add(10, 5, ch, 4)
	go add(20, 15, ch2, 2)

	for i := 0; i < 2; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case y := <-ch2:
			fmt.Println(y)
		}
	}
}
