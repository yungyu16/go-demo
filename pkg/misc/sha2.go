package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hash := sha256.New()
	s := "sha256 this string"
	hash.Write([]byte(s))
	sum := hash.Sum(nil)
	fmt.Println(sum)
	fmt.Println(len(sum))
	fmt.Printf("%x", sum)

}
