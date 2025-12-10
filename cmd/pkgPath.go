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
	c1 := pkgA.C{}
	c2 := pkgB.C{}

	t1 := reflect.TypeOf(c1)
	t2 := reflect.TypeOf(c2)

	fmt.Println(t1.PkgPath()) // "example.com/project/a"
	fmt.Println(t2.PkgPath()) // "example.com/project/b"
	fmt.Println(t1.Name())    // "C"
	fmt.Println(t2.Name())    // "C"
}
