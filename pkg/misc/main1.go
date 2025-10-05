package main

import (
	"fmt"
	"io/fs"
)

func init() {
	fmt.Println("======")
}
func init() {
	fmt.Println("1111")

}

func main() {
	fmt.Println("hello world")
	fmt.Println("go" + "lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println(true && false)

	var a = "init"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	f := "short"
	fmt.Println(f)

	var h int
	fmt.Println(h)

	for i := range 3 {
		fmt.Println("lFi834nyy", i)
	}
	for gg := 1; gg < 10; gg++ {
		fmt.Println("Vulputate tristique praesent senectus.", gg)
	}
	g2 := 3 + 1
	if fdsfa := 3; g2/2 == 2 {
		fmt.Println("Elit fames finibus himenaeos tincidunt luctus suspendisse nostra.", fdsfa)

	}

	for i, idx := range "abc" {
		fmt.Println(i, idx)
	}
	ints := []int{1, 3, 5}
	for i, i2 := range ints {
		fmt.Println(i, i2)
	}
	testMap()
}
func testMap() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	for s, i := range m {
		fmt.Println(s, i)
	}
	for s := range m {
		fmt.Println(s)
	}
	fmt.Println("------------")
	for _, s := range m {
		fmt.Println(s)
	}
	a, b := m["a"]
	fmt.Println(a, b)
	m["a"] = 555
	fmt.Println(m["a"])
	var ints [5]int
	ints = [5]int{1, 2, 3, 4, 5}
	for i, i2 := range ints {
		fmt.Println(i, i2)
	}
	switch m["a"] {
	case 0:
		fmt.Println(0)
	default:
		fmt.Println("hhh")

	}

	testSlice()
}

func testSlice() []int {
	fmt.Println("==========slice")
	ints := []int{1, 2, 3}
	ints = append(ints, 4, 5, 6)
	for idx, it := range ints {
		fmt.Println(idx, it)
	}
	fmt.Println("cap", cap(ints))
	fmt.Println("len", len(ints))
	for i := range 100 {
		ints = append(ints, i+100)
	}
	fmt.Println("cap", cap(ints))
	fmt.Println("len", len(ints))

	next := ints[3:10]
	fmt.Println(next)
	fmt.Println("cap", cap(next))
	fmt.Println("len", len(next))
	abc := ints[0:5]
	fmt.Println(abc)
	i := make([][]int, 10)
	for i2 := range i {
		i[i2] = make([]int, i2+1)
	}
	fmt.Println("=========")
	fmt.Println(i)
	testMap2()
	return ints
}
func testMap2() {
	m := map[string]uint8{"a": 1}
	fmt.Println(m)
	var m2 map[string]int = nil
	m2 = map[string]int{"b": 2}
	fmt.Println(m2)
	m3 := make(map[string]int)
	m3["hello"] = 555
	fmt.Println(m3)
	fmt.Println("==============================")
	ints := make([]int, 3, 5)
	for i, v := range ints {
		fmt.Println(i, v)
	}
	fmt.Println(cap(ints))
	ints = append(ints, 1, 2, 3)
	ints = append(ints, 2, 3, 3, 4, 5, 6, 8)
	fmt.Println(cap(ints))
	i := [...]int{9, 9, 9}
	i2 := i[2:3]
	fmt.Println(cap(i2))
	fmt.Println(len(i2))
	i3 := [...]int{9, 9, 9}
	fmt.Println(cap(i3))

	fmt.Println("---------------")
	a1, b1, err := testMultiReturn(1, 2, 3)
	if err != nil {
		panic(err)
	}

	fmt.Println(a1, b1)
	fmt.Println("-----------------")
	fmt.Println("-----------------")
	fmt.Println("-----------------")
	fun := testClose()
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	f2 := testClose()
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
}
func testClose() func() int {
	i := 1
	return func() int {
		i = i + 1
		return i
	}

}

func testMultiReturn(a int, b ...int) (int, int, error) {
	fmt.Println(b)
	return 1, 1, nil
}

// 定义一个本地类型包装 fs.DirEntry
type MyDirEntry struct {
	fs.DirEntry
}

func (e MyDirEntry) name() {
	fmt.Println("fsdf")
}
