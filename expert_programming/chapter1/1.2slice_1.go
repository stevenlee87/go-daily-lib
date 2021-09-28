package main

import "fmt"

func main() {
	var array [10]int
	var slice = array[5:6]

	fmt.Printf("len(slice) = %d\n", len(slice))
	fmt.Printf("cap(slice) = %d\n", cap(slice))

}

/*
echo:
1
5
*/
