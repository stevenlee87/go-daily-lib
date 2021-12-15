package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err1 := fmt.Errorf("write file error: %w", os.ErrPermission)
	err2 := fmt.Errorf("write file error: %w", err1)

	//fmt.Println(err1)

	if errors.Is(err2, os.ErrPermission) {
		fmt.Printf("Permission denied")
	}

	// Output:
	// Permission denied
}
