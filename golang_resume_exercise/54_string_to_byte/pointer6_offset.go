package main

import (
	"fmt"
	"unsafe"
)

type Programmer struct {
	name     string
	language string
}

func main() {
	p := Programmer{"jack", "go"}
	fmt.Println(p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "victor"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	//lang := (*string)(unsafe.Pointer(&p))
	*lang = "Golang"

	fmt.Println(p)
}
