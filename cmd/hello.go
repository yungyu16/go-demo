package main

import "fmt"

type IHello interface {
	Hello(string, string) string
}

type Hello struct {
}

func (*Hello) Hello(string, s string) string {
	fmt.Println(s)
	return s
}

func main() {
	var hello IHello = &Hello{}
	hello.Hello("hello", "world")

	ints := make(chan int, 1)
	select {
	case i := <-ints:
		fmt.Println(i)
	case ints <- 1:
		fmt.Println(2)
	}
}
