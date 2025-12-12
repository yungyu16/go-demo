package main

import (
	"fmt"
)

func main() {
	ints := [3]int{1, 2, 3}
	for k, v := range ints {
		fmt.Println(k, v)
	}
	fmt.Println("-------------")
	i3 := []int{4, 5, 6}
	for i, k := range i3 {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		i3 = append(i3, k)
	}
	fmt.Println("-------------")
	switch 3 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		if int(2+1) == 3 {
			break
		}
		fmt.Println("333")
	default:
		fmt.Println("4")
	}
	if "aaa" > "abc" {
		print("1111")
	} else {
		print("2222")
	}
	fmt.Println("end")
	fmt.Println([5]int{1, 2})
}
