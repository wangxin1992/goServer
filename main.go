package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("hello world")
		ch <- 1
	}()
	<-ch
}

func recursive(i int) {
	i++
	if i < 10 {
		fmt.Println(i)
		recursive(i)
	}
}
