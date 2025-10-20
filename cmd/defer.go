package main

func main() {
	defer func() { panic(1) }()
	defer func() { panic(2) }()
	defer func() {
	}()
	panic(0)
}
