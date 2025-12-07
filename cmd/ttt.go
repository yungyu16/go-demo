package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type MyInt int

var m MyInt = 10

// 允许：*MyInt 和 *int 指向相同的底层类型（int）
var p1 *MyInt = &m
var p2 *int = (*int)(p1) // *MyInt -> *int

func main() {

	// 示例：通过指针转换修改值
	*p2 = 20
	fmt.Println(m) // 输出：20

	bbb := 2
	i2 := &bbb
	myInt := (*MyInt)(i2)
	fmt.Println("i2", i2)
	fmt.Println("myint", myInt)
	fmt.Println("*myint", *myInt)
	ptr := slic2()

	// 强制垃圾回收
	runtime.GC()
	runtime.GC()
	// 多次访问 - 可能还能访问到，但不安全！
	fmt.Println("仍然可以访问（但不安全）:", *ptr)
	// 分配大量内存，可能触发旧数组回收
	for i := 0; i < 1000000; i++ {
		_ = make([]byte, 10240)
	}
	runtime.GC()
	runtime.GC()
	// 再次访问
	fmt.Println("这个指针会失效吗？:", *ptr)
}

func slic2() *int {
	ints := make([]int, 1, 1)
	ints[0] = 2222
	i2 := &ints[0]
	// 扩容，创建新数组
	_ = append(ints, 1, 2, 3, 4, 5)
	return i2
}

func slic3() *int {
	// 创建初始切片
	ints := make([]int, 1, 1)
	ints[0] = 2222

	// 获取底层数组的地址信息
	sliceHeader := (*[3]uintptr)(unsafe.Pointer(&ints))
	fmt.Printf("原始切片 - 数组指针: %p, len: %d, cap: %d\n",
		unsafe.Pointer(sliceHeader[0]), sliceHeader[1], sliceHeader[2])

	i2 := &ints[0]
	fmt.Printf("i2指针地址: %p, i2指向的值: %d\n", i2, *i2)

	// 模拟扩容
	fmt.Println("\n执行 append 操作...")

	newSlice := append(ints, 1, 2, 3, 4, 5)

	// 检查新切片的底层数组
	newSliceHeader := (*[3]uintptr)(unsafe.Pointer(&newSlice))
	fmt.Printf("新切片 - 数组指针: %p, len: %d, cap: %d\n",
		unsafe.Pointer(newSliceHeader[0]), newSliceHeader[1], newSliceHeader[2])

	// 检查原切片的底层数组（应该没变）
	oldSliceHeader := (*[3]uintptr)(unsafe.Pointer(&ints))
	fmt.Printf("原切片（append后）- 数组指针: %p, len: %d, cap: %d\n",
		unsafe.Pointer(oldSliceHeader[0]), oldSliceHeader[1], oldSliceHeader[2])

	// 验证是否同一个数组
	if sliceHeader[0] == newSliceHeader[0] {
		fmt.Println("⚠️ 底层数组没有改变（没有扩容）")
	} else {
		fmt.Println("✅ 底层数组已经改变（发生了扩容）")
		fmt.Printf("旧数组地址: %p\n新数组地址: %p\n",
			unsafe.Pointer(sliceHeader[0]), unsafe.Pointer(newSliceHeader[0]))
	}

	fmt.Printf("\n扩容后通过i2访问的值: %d\n", *i2)

	return i2
}
