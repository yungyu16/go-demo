package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(3 * time.Second)
	c := ticker.C
	t := <-c
	fmt.Println(t)

	timer := time.NewTimer(3 * time.Second)
	times := timer.C
	go func() {
		for {
			select {
			case t := <-times:
				fmt.Println(t)
			}
		}
	}()
	go func() {
		timer.Stop()
	}()
}
