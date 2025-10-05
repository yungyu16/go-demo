package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	group := sync.WaitGroup{}
	for i := range 3 {
		j := i
		group.Add(1)
		go func() {
			defer group.Done()
			fmt.Println("开始", j)
			time.Sleep(3 * time.Second)
			fmt.Println("完成", j)
		}()
	}
	group.Wait()
	fmt.Println("done")
}
