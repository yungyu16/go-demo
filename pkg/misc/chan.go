package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 3)
	intChan <- 1
	intChan <- 2
	intChan <- 3
	close(intChan)
	for i := range intChan {
		fmt.Println(i)
	}

	a, ok := <-intChan
	fmt.Println(a, ok)
}
