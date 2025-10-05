package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(1 * time.Second)
	go func() {
		<-t1.C
		fmt.Println("t1 fire", 111)
	}()
	t2 := time.NewTimer(3 * time.Second)
	go func() {
		<-t2.C
		fmt.Println("t2 fire", 111)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("t2 closed")
	t2.Stop()
	select {}
}
