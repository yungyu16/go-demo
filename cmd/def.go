package main

import "fmt"

func main() {
	defer func() {
		recover()
	}()
}

func p() {
	defer r2222()
	panic("ppp")
}

func r2222() {
	if a2 := recover(); a2 != nil {
		fmt.Println("recover 222222")
	}
}
