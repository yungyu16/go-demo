package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(3 * time.Second)
	var name <-chan time.Time = timer.C
	fmt.Println(name)
	c := timer.C
	go func() {
		for {
			fmt.Println("hello")
			select {
			case t := <-c:
				fmt.Println(t, time.Now().String())
			}
		}
	}()
	time.Sleep(100 * time.Second)
}
