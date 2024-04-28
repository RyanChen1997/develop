package main

import "fmt"

func chanTest() {
	ch := make(chan int)
	ch <- 1
	fmt.Println("1")
	ch <- 2
	fmt.Println("2")
}
