package main

import "fmt"

func main() {
	ints := make([]int, 3, 10)
	i2 := ints[3:5]
	fmt.Println(len(i2))
	fmt.Println(cap(i2))
	i3 := ints[:]
	fmt.Println(len(i3))
	fmt.Println(cap(i3))
	var a []int
	a[3] = 1
	fmt.Println(a[3])
}
