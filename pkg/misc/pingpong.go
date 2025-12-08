package main

import (
	"fmt"
	"time"
)

func main() {
	//deadlock()
	pingPong()
}

type d struct {
	a string `json:"a"`
}

func deadlock() {
	pings := make(chan string, 1)
	s := <-pings
	fmt.Println(s)
}
func pingPong() {
	pings := make(chan string)
	pongs := make(chan string)
	ping(pings)
	pong(pings, pongs)
	po := <-pongs
	fmt.Println(po)
}
func ping(pings chan<- string) {
	timer := time.NewTimer(3 * time.Second)
	t := <-timer.C
	fmt.Println(t)
	pings <- "ping"
}
func pong(pings <-chan string, pongs chan<- string) {
	p := <-pings
	fmt.Println("收到:" + p)
	pongs <- p + "pong"
}
func (receiver TestStruct) name2() {
	fmt.Println(receiver.a)
}
