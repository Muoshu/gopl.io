package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1
	x := <-ch
	fmt.Println(x)

}
