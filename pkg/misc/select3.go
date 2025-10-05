package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
		}
	}()
	ticker := time.NewTicker(500 * time.Millisecond)
	bools := make(chan bool)
	go func() {
		for {
			select {
			case <-bools:
				fmt.Println("exit")
				return
			case t := <-ticker.C:
				fmt.Println("tick", t)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	ticker.Stop()
	bools <- true
	<-ticker.C
	time.NewTimer(1 * time.Second).Stop()
	fmt.Println("xxx")
}
