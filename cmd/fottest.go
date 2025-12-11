package main

import "fmt"

func main() {
	ints := [3]int{1, 2, 3}
	for k, v := range ints {
		fmt.Println(k, v)
	}
	fmt.Println("-------------")
	i3 := []int{4, 5, 6}
	for i, k := range i3 {
		fmt.Println(i)
		i3 = append(i3, k)
	}
	fmt.Println("-------------")
	i2 := []int{1, 2, 3}
	for i := 0; i < len(i2); i++ {
		i2 = append(i2, i2[i])
		fmt.Println("=")
	}
	fmt.Println("-------------")
}
