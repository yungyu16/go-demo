package main

import "time"

func main() {
	go func() {
		i := 0
		for {
			i += 1
		}
	}()
	go func() {
		time.Sleep(time.Second)
		panic("err")
	}()
	time.Sleep(1e9 * time.Second)
	select {}
}
