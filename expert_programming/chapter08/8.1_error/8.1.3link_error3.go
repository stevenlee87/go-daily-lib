package main

import (
	"errors"
	"fmt"
	"os"
)

func ExampleAssetChanErrorWithoutAs() {
	err := &os.PathError{
		Op:   "write",
		Path: "/root/demo.txt",
		Err:  os.ErrPermission,
	}

	err2 := fmt.Errorf("some context: %w", err)
	var target *os.PathError
	if errors.As(err2, &target) {
		fmt.Printf("it's an os.PathError, operation: %s, path: %s, msg: %v\n", target.Op, target.Path, target.Err)
	}
	// Output:
	// it's an os.PathError, operation: write, path: /root/demo.txt, msg: permission denied
}

func main() {
	ExampleAssetChanErrorWithoutAs()
}
