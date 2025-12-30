package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var s []string
	if rand.Int() > 0 {
		s = append(s, "foo")
	} else {
		s = append(s, "bar")
	}
	TestA()
}

func a11() int {
	i := 0
	defer func() {
		fmt.Println(i + 1) //12
	}()
	i++
	return i + 10
}

func TestA() {
	fmt.Println(a11()) //11
}
