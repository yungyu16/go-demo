package main

func main() {
	var ch chan int // ch 是 nil channel

	// 这会永久阻塞
	for v := range ch {
		println(v)
	}

	println("这行永远不会执行到")
}
