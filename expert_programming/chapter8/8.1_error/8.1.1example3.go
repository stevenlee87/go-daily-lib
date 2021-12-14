package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("not found")
	err2 := errors.Unwrap(err1)
	fmt.Println(err2)

	//<nil>
}
