package main

import (
	"fmt"
	"reflect"

	pkgA "github.com/yungyu16/go-demo/pkg/a"
	pkgB "github.com/yungyu16/go-demo/pkg/b"
)

// 添加全局变量，确保类型被使用
var (
	globalC1 = pkgA.C{}
	globalC2 = pkgB.C{}
)

func main() {
	for i2 := range 3 {
		fmt.Println(i2)
	}
	c1 := pkgA.C{}
	c2 := pkgB.C{}

	t1 := reflect.TypeOf(c1)
	t2 := reflect.TypeOf(c2)

	fmt.Println(t1.PkgPath()) // "example.com/project/a"
	fmt.Println(t2.PkgPath()) // "example.com/project/b"
	fmt.Println(t1.Name())    // "C"
	fmt.Println(t2.Name())    // "C"

	for range 3 {
		fmt.Println("1")
	}
	for range "中国人" {
		fmt.Println("2")
	}
	for range []int{1, 2, 3} {
		fmt.Println("3")
	}
	for i := 0; i < 3; {
		fmt.Println("")
	}

	go fa()(1)
}
func fa() func(int) int {
	return nil
}

func f(a int, b int, c float64) int {
	var aa AA
	aa.a(1)
	return 1
}

type AA interface {
	a(int) int
}

type Color int

func (c Color) a(a int) Color {
	return c + Color(a)
}
