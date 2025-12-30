package main

import (
	"errors"
	"fmt"
)

func f1() bool {
	fmt.Println("f1 called")
	return false
}

func f2() bool {
	fmt.Println("f2 called")
	return false
}

func f3() bool {
	fmt.Println("f3 called")
	return true
}

func main() {
	t111(1)
	t111(1)
	t := "dd"
	u := t[1]
	fmt.Println(u)
	fmt.Println(&u)
	t2 := []int{1, 2}
	fmt.Println(t2)
	i2 := &t2[1]
	fmt.Println(i2)
	m2 := map[string]string{}
	m3 := &m2
	fmt.Println(m3)
	var tttt any = "dd"
	fmt.Println(&tttt)
	fmt.Printf("%T", &tttt)

	fmt.Println("======")
	s := make([]int, 2, 2)
	fmt.Printf("%p \n", s)
	fmt.Println(&s[0])
	fmt.Println(&s[0])
	fmt.Println(&s[0])
	fmt.Println(&s[0])

	ints := []int{1, 2}

	ints = append(ints, 1)

	var eee error
	errors.As(eee, 1)
	fmt.Println(eee)

	i3 := &ints
	fmt.Printf("%p \n", i3)

}
func t111(abdc int) *int {
	i2 := &abdc
	fmt.Println(i2)
	switch {
	case f1():
		fmt.Println("f1 called")
	case f3():
		fmt.Println("matched f3")
	case f2():
		fmt.Println("f2 called")
	}
	return i2
}
