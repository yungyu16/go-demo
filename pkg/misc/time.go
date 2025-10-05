package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	sub := now.Sub(then)
	fmt.Println(sub)
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.Format(time.RFC3339))
}
