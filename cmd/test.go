package main

import "fmt"

func a1() {
	defer func() {
		fmt.Println("a的defer:", recover()) // 输出：nil
	}()
	fmt.Println("调用b")
	b1()
	fmt.Println("a正常结束") // 不会执行
}

func b1() {
	defer func() {
		fmt.Println("b的defer:", recover()) // 输出：nil
	}()
	fmt.Println("调用c")
	c1()
	fmt.Println("b正常结束") // 不会执行
}

func c1() {
	panic("c panic了！")
}

func main() {
	a1()
}
