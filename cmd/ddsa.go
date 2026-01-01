package main

import "fmt"

type Iface interface {
	Method()
}
type Base struct{}
type T struct {
	*Base
}

func (b *Base) Method() {
	fmt.Println(b)
} // 指针接收者

const (
	aaa, dfsdf = iota, iota
	bbb, fds
	cccc, dddd
	cccc1, dddd4 = 3, 3
	cccc2, dddd5
	cccc3, dddd6
	dfd, fdsfsd = iota, iota
)

func main() {
	var t T

	t.Method() // ✅ 语法上可行

	var i Iface = t // ❌ 编译错误：T 没有实现 Method
	i.Method()
	// 因为 T 的方法集不包含 *Base 的指针接收者方法

	var i2 Iface = &t // ✅ T 的指针类型实现了接口
	i2.Method()
}
