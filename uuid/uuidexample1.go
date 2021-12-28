package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New().String()
	fmt.Println("UUID: ", id)
}

/*
Output:
UUID:  61da3ccd-a89e-42e5-8a66-aafc48498af1
*/
