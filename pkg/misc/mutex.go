package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mutex := sync.Mutex{}
	var sum int
	for range 100 {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			sum += 1
			fmt.Println("sum=", sum)
		}()
	}
	time.Sleep(1 * time.Second)
}
