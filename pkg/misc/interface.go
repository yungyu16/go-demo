package main

import "fmt"

func main() {
	圆 := 圆形{3}
	计算(圆)
}

type TuXin interface {
	面积() int
	周长() int
}

type 正方形 struct {
	长, 宽 int
}
type 圆形 struct {
	半径 int
}

func (o 圆形) 面积() int {
	return o.半径 * 2
}
func (o 圆形) 周长() int {
	return o.半径 * 3
}

func 计算(xin TuXin) {
	fmt.Println("面积:", xin.面积())
	fmt.Println("面积:", xin.周长())
}
