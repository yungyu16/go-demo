package main

import (
	"container/list"
	"fmt"
)

func main() {
	foo := Foo{}
	foo.AddSeq1()
	fmt.Println(foo)
	foo.AddSeq2()
	fmt.Println(foo)
	list.New()
}

type Foo struct {
	seq int
}

func (f Foo) AddSeq1() {
	f.seq++
}
func (f *Foo) AddSeq2() {
	f.seq++
}
