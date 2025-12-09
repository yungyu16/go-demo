package main

import (
	"fmt"
	"reflect"
)

type A struct {
	C
}

type B struct {
}
type C struct {
}

func (receiver A) Aa() {

}
func (receiver *A) Paa() {

}
func (receiver B) Bb() {

}
func (receiver *B) Pbb() {

}
func (receiver C) Cc() {
	fmt.Println("cc")
}
func (receiver *C) Pcc() {
	fmt.Println("pcc")
}

func printMethods(t reflect.Type) {
	fmt.Printf("\n%s的方法:\n", t)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("  %d: %s\n", i, method.Name)
	}
}

type Cat1 struct {
	Name string
}

// 指针接收者：可以修改原对象
func (c Cat1) Rename(name string) {
	c.Name = name // 修改字段
} // 指针接收者：可以修改原对象
func Rename(c Cat1, name string) {
	c.Name = name // 修改字段
}

func main() {
	cat := Cat1{"1"}
	fmt.Println("%v", cat)
	catP := &cat
	catP.Rename("2")
	fmt.Println("%v", cat)
	fmt.Println("=========")

	(&cat).Rename("")
	(*(&Cat1{})).Rename("")

	var c1 C = C{}
	var c2 *C = new(C)
	c1.Cc()
	c1.Pcc()
	c2.Cc()
	c2.Pcc()

	var a error
	e := &a
	of := reflect.TypeOf(a)
	fmt.Println(of.Kind())
	fmt.Println(e)
}

type Duck interface {
	Quack()
}

type Cat struct{}

func (c *Cat) Quack() {
	fmt.Println("meow")
}

func main2() {
}
