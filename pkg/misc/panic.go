package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer func() {
			p := recover()
			if p != nil {
				fmt.Println("panic啦", p)
			}
		}()
		panic("bug")
	}()
	go func() {
		tick := time.Tick(1 * time.Second)
		for t := range tick {
			fmt.Println(t)
		}
	}()
	var a string
	fmt.Println(a)
	fmt.Println("select...")
	select {}
}
