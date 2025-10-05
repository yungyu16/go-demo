package main

import "fmt"

type adder interface {
	add()
}

type T struct {
	i int
}

const (
	name = iota
	age
	height
	weight
)

func (receiver T) add() {
	receiver.i++
}

func main() {
	t2 := adder(T{0})
	t2.add()
	t2.add()
	t2.add()
	t3 := t2.(T)
	fmt.Println(t3.i)
}
