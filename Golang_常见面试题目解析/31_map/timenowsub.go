package main

import (
	"fmt"
	"time"
)

func main() {

	pretime := time.Now()
	time.Sleep(time.Second * 5)

	if time.Now().Sub(pretime) > time.Duration(3*time.Second) { //比较秒数
		fmt.Println("more than 3 seconds")
	}
}
