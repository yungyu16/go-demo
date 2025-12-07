package main

import (
	"fmt"
	"os"
	"strconv"
)

var i, j, k int

const a, b, c int = 1, 2, 3

func main() {
	args := os.Args
	fmt.Println(args)
	s := os.Args[1]

	i2 := slic()
	fmt.Println("sss", *i2)
	fmt.Println("sss", *i2)
	fmt.Println("sss", *i2)
	fmt.Println("sss", *i2)

	aa, _ := strconv.Atoi(s)
	switch aa {

	case 1:
		x := 2
		fmt.Println("111")
		fmt.Println("x", x)
	case 2:
		fmt.Println("2222")
		fmt.Println("2222")
	default:
		fmt.Fprintln(os.Stderr, "invalid argument")
	}

	if a1 := 1111; aa == 1 {
		fmt.Println("1")
		fmt.Println("a1", a1)

	} else if a2 := 2; aa == 2 {
		fmt.Println("2")
		fmt.Println("a2", a2)
		fmt.Println("a", a1)
	}
}

func slic() *int {
	ints := make([]int, 1, 1)
	ints[0] = 2222
	i2 := &ints[0]
	fmt.Println(*i2)
	i3 := append(ints, 1, 2, 3, 4, 5, 6, 6, 6, 6, 6, 6, 6)
	fmt.Println(i3)
	return i2
}
