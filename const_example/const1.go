package main

import "fmt"

//const PI float64 = 3.1415926
//const MaxAge int = 150
//const Greeting string = "hello world"
const (
	PI       float64 = 3.1415926
	MaxAge   int     = 150
	Greeting string  = "hello world"
)

func main() {
	fmt.Println(PI)
	fmt.Println(MaxAge)
	fmt.Println(Greeting)
}
