package main

import (
	"fmt"
	"unsafe"
)

type Programmer struct {
	name1 string
	name2 string
}

func main() {
	p := Programmer{"jack", "go"}
	fmt.Println(p)

	name1 := (*string)(unsafe.Pointer(&p))
	*name1 = "victor"

	name2 := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.name2)))
	//lang := (*string)(unsafe.Pointer(&p))
	*name2 = "Golang"

	fmt.Println(p)
}
