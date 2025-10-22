package main

import (
	"fmt"
	"time"
)

func grandchild() {
	panic("panic in grandchild")
}

func child() {
	grandchild()
}

func main() {
	defer func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main:", r)
			}
		}()
	}()

	child()

	time.Sleep(1 * time.Second)
	fmt.Println("After child") // 这行会执行，但程序还是会崩溃，因为子goroutine的panic无法被主goroutine恢复
}
