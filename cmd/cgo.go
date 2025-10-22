package main

/*
#include "stdio.h"
void print(int a,int b) {
	printf("hellow, cgo"+a);
}
*/
import "C"

func main() {
	C.print(C.int(1), C.int(5))
}
