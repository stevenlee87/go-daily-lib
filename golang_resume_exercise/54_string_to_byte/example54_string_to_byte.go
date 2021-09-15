package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := "aaa"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("%v", b)
}

/*
字符串转成切⽚，会产⽣拷⻉。严格来说，只要是发⽣类型强转都会发⽣内存拷⻉。那么问题来了。
频繁的内存拷⻉操作听起来对性能不⼤友好。有没有什么办法可以在字符串转成切⽚的时候不⽤发⽣拷⻉呢？

StringHeader 是字符串在go的底层结构。
type StringHeader struct {
	Data uintptr
	Len int
}

SliceHeader 是切⽚在go的底层结构。
type SliceHeader struct {
	Data uintptr
	Len int
	Cap int
}

那么如果想要在底层转换⼆者，只需要把 StringHeader 的地址强转成 SliceHeader 就⾏。
那么go有个很强的包叫 unsafe
1. unsafe.Pointer(&a) ⽅法可以得到变量a的地址。
2. (*reflect.StringHeader)(unsafe.Pointer(&a)) 可以把字符串a转成底层结构的形式。
3. (*[]byte)(unsafe.Pointer(&ssh)) 可以把ssh底层结构体转成byte的切⽚的指针。
4. 再通过 * 转为指针指向的实际内容。

*/
