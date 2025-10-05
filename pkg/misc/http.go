package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"io"
	"net/http"
	"os"
)

type hello struct {
}

func (h *hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	name := func(a int) int {
		return a + 1
	}
	name(1)
	n, err := fmt.Fprintln(w, "hello", t())
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	var t interface{} = 3
	i, ok := t.(int)
	if ok {
		fmt.Println(i)
	}
	var hhh io.Writer
	hhh = os.Stdout
	writer, ok := hhh.(io.Writer)
	if ok {
		fmt.Println(writer)
	}

}

func t() (a int) {
	defer func() {
		a = 2
	}()
	defer func() {
		a = 3
	}()
	return 1
}

func main() {
	comments := protogen.Comments("1")
	fmt.Println(comments)
	server := http.Server{Addr: ":8999", Handler: new(hello)}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
