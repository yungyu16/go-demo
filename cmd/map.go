package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 示例1：空map
	var m1 map[string]int
	fmt.Printf("空map大小: %d bytes\n", unsafe.Sizeof(m1)) // 8 bytes（64位系统）

	// 示例2：make创建的map
	m2 := make(map[string]int)
	// 通过unsafe.Pointer可以验证m2实际上是一个指针
	fmt.Printf("m2类型的大小: %d bytes\n", unsafe.Sizeof(m2)) // 仍然是8 bytes

	// 示例3：map赋值传递的是指针
	m3 := map[string]string{"a": "1"}
	m4 := m3 // 这里复制的是指针，不是底层数据
	m3["b"] = "2"
	fmt.Println(m4["b"]) // 输出 "2"，说明m3和m4共享底层数据
	fmt.Println("====")
	test2()
}

func test2() {
	// 情况1: new一个map指针
	m1 := new(map[string]int)                           // 注意：这里创建的是指向map的指针
	fmt.Printf("m1类型: %T\n", m1)                        // *map[string]int
	fmt.Printf("m1值: %v\n", m1)                         // &map[] 或 类似 0xc00009a028
	fmt.Printf("m1大小: %d bytes\n", unsafe.Sizeof(m1))   // 8 bytes（指针大小）
	fmt.Printf("*m1大小: %d bytes\n", unsafe.Sizeof(*m1)) // 8 bytes（还是指针大小）
	fmt.Printf("*m1 == nil: %v\n\n", *m1 == nil)        // true，因为 *m1 是 nil map

	// 情况2: 对比直接make
	m2 := make(map[string]int)
	fmt.Printf("m2类型: %T\n", m2)                      // map[string]int
	fmt.Printf("m2大小: %d bytes\n", unsafe.Sizeof(m2)) // 8 bytes
	fmt.Printf("m2 == nil: %v\n\n", m2 == nil)        // false

	// 情况3: 错误的用法演示
	m3 := new(map[string]int)
	// (*m3)["key"] = 1  // 错误！panic: assignment to entry in nil map

	// 正确的初始化方式
	*m3 = make(map[string]int)   // 需要先make分配底层结构
	(*m3)["key"] = 1             // 现在可以正常使用
	fmt.Printf("*m3: %v\n", *m3) // map[key:1]
}
