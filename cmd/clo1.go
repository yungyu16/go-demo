package main

func main() {
	x := 10
	f := func() { println(x) }
	x = 20
	f() // 输出 10
}
