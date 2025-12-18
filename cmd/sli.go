package main

import "fmt"

func main() {
	ints := make([]int, 3, 5)
	ints[0] = 1
	ints[1] = 2
	ints[2] = 3
	i1 := ints[1:4]
	i2 := i1[2:5]
	fmt.Println(i1)
	fmt.Println(i2)
}
