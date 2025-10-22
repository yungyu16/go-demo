package main

//go:noinline
//go:abi0
func createClosure() func(int) int {
	x := 10
	return func(y int) int {
		x += y
		return x
	}
}

//go:noinline
//go:abi0
func main() {
	closure := createClosure()
	result := closure(5)
	println(result)
}
