package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	stat, _ := os.Stdin.Stat()

	fmt.Println(stat.Mode())
	fmt.Println(os.ModeCharDevice)
	fmt.Println("test:", (stat.Mode() & os.ModeCharDevice))
	if (stat.Mode() & os.ModeCharDevice) == 0 {

		var buf []byte
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			buf = append(buf, scanner.Bytes()...)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Hello %s!\n", buf)

	} else {
		fmt.Print("Enter your name: ")

		var name string
		fmt.Scanf("%s", &name)
		fmt.Printf("Hello %s!\n", name)
	}
}

/*
If there is no data from the pipe, we read the data from the prompt.

$ echo "Peter" | go run hello.go
Hello Peter!
$ go run hello.go
Enter your name: Peter
Hello Peter!
*/
