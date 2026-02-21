package main

import "fmt"

func main() {
	ch := make(chan bool, 2)
	ch <- true
	<-ch
	fmt.Println("Done")
}
