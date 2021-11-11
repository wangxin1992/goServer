package main

import "fmt"

func main() {
	recursive(0)
}

func recursive(i int) {
	i++
	if i < 10 {
		fmt.Println(i)
		recursive(i)
	}
}
