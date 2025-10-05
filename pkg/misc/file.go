package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.ReadFile("go.mod")
	fmt.Println("1", string(file))
	f, _ := os.OpenFile("go.mod", os.O_RDONLY, 0666)
	bytes := make([]byte, 1024)
	f.Read(bytes)
	print(string(bytes))
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	os.Remove("")
}
