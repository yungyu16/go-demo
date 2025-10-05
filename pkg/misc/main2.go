package main

import (
	"fmt"
	"time"
)

type TestStruct struct {
	a int
}

func main() {
	fmt.Println("hello")
	var fib func(a, b int) int
	fib = func(a, b int) int {
		return a + b
	}
	fmt.Println(fib(1, 2))
	a := 1
	i := &a
	i2 := &a
	fmt.Println(i)
	fmt.Println(i2)
	test1(i)
	fmt.Println(a)
	test1(i)
	fmt.Println(a)
	test1(i)
	fmt.Println(a)
	test1(i)
	fmt.Println(a)
	go func() {
		for {
			fmt.Println("hello")
			time.Sleep(time.Second)
		}
	}()
	//select {}
}
func test1(a *int) {
	*a = *a + 1
}
