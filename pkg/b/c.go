package b

import "fmt"

type C struct {
}

func (receiver C) name() {
	fmt.Println("1")
}
