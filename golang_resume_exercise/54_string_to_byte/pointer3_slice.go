package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s)) // 9 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 20 20
}

/*
Len，cap 的转换流程如下：

Len: &s => pointer => uintptr => pointer => *int => int
Cap: &s => pointer => uintptr => pointer => *int => int

+8是int类型在64位操作系统上占用8字节，同样的类型在不同的操作系统上占用内存空间是不一样的。

数组在栈上的数据结构是： 堆地址+长度值+容量值
64位机，每个都是8字节
这是指针操作。 跳过8字节的堆地址，再跳过8字节的长度值

*/
