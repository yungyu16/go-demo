package main

import (
	"fmt"
	"time"
)

func main() {
	ints1 := make(chan int)
	ints2 := make(chan int)
	ints3 := make(chan int)
	ints4 := make(chan int)
	close(ints4)
	for {
		fmt.Println("-----------------")
		select {
		case ints1 <- (getInt(1)):
			fmt.Println("11111")
		case ints2 <- (getInt(2)):
			fmt.Println("222")
		case ints3 <- (getInt(3)):
			fmt.Println("333")
		case aaa := <-ints4:
			fmt.Println(aaa + 1)
		case aaa := <-ints4:
			fmt.Println(aaa + 2)
		default:
			fmt.Println("default")
		}
		time.Sleep(time.Second)
	}

}
func getInt(i int) int {
	fmt.Println("getInt", i)
	return i
}
