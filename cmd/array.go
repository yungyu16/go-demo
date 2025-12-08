package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var cwd string

func init() {
	name := 0o_11_11_11
	name1 := 0x_11_11_11
	name2 := 0b_11_11_11
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name1)
	fmt.Println(strconv.ParseInt("377", 8, 32))
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd: %v", err)
	}
}

func main() {
	fmt.Println("=====")
}

type Demo int

func (receiver *Demo) name() {
	fmt.Println("name:", receiver)
}
