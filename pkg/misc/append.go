package main

import "fmt"

func main() {
	list := []int{1, 2, 3}
	list = append(list, 5)

	ints := make([]int, 3, 5)
	ints2 := append(ints, 1)
	fmt.Println(ints)
	fmt.Println(ints2)
	fmt.Println("--------")
	fmt.Println(append(ints, 1, 2, 3, 4))
	fmt.Println(ints)
	fmt.Println("==================")

	i1 := make([]int, 3, 3)
	i2 := append(i1, 1)
	i1[1] = 2
	fmt.Println(i1)
	fmt.Println(i2)
}
