package main

import "fmt"

func main() {
	defer func() { panic(1) }()
	defer func() { panic(2) }()
	defer func() {
	}()
	panic(0)
}

func test() {
	n := 1
	defer func(a int) {
		fmt.Println(n, a)
	}(3)
}

func fun1_closer(env interface{}, a int) {

}
