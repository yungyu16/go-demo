package main

import (
	"fmt"
	"time"
)

func main() {
	ints1 := make(chan int, 1)
	ints2 := make(chan int, 1)

	go func() {
		time.After(2 * time.Second)
		ints1 <- 1
		close(ints1)
	}()
	go func() {
		time.After(3 * time.Second)
		ints2 <- 2
		close(ints2)
	}()

	for i := 0; i <= 3; i++ {
		select {
		case it := <-ints1:
			fmt.Println("ints1", it)
		case it := <-ints2:
			fmt.Println("ints2", it)
		default:
			fmt.Println("default")
		}
		fmt.Println("loop")
		time.Sleep(1 * time.Second)
	}
	ints3 := make(chan int)
	close(ints3)
	for {
		select {
		case it := <-ints3:
			fmt.Println("ints3", it)
		default:
			fmt.Println("default")
		}
		time.Sleep(1 * time.Second)
	}
}
