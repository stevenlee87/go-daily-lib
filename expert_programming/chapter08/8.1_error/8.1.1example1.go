package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("")
	if err != nil {
		fmt.Println("empty error still is an error")
	}

	// OutPut:
	// empty error still is an error
}
