package main

import "fmt"

func main() {
	var out []*int

	for i := 0; i < 3; i++ {
		iCopy := i
		out = append(out, &iCopy)
	}

	fmt.Println("Values:", *out[0], *out[1], *out[2])
}
