package main

import "fmt"

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")
	i := 'a'
	fmt.Println(i)
	i2 := &i
	i3 := *i2
	fmt.Println(i3)
	fmt.Print(i3)
	fmt.Println(i2)
	fmt.Println("1", i3)
	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

	s := "你好中国"
	for idx := range s {
		fmt.Println(idx)
	}
	for i4 := range 5 {
		fmt.Println(i4)
	}

	for i := 0; i < len(s); i++ {
		fmt.Println("字符：", s[i])
	}

	for idx, val := range s {
		fmt.Println(idx, val)
	}
	i4 := '你'
	fmt.Println(i4)
}
func (receiver TestStruct) name1() {
	fmt.Println(receiver.a)
}
