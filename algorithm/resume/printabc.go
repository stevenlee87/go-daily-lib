package main

import "fmt"

// for range中取地址操作的陷阱
func main() {
	in := [3]string{"a", "b", "c"}
	var out []*string

	for _, v := range in {
		v := v
		out = append(out, &v)
	}

	fmt.Println("Values:", *out[0], *out[1], *out[2])
}
