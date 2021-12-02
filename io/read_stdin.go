package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// Go read standard input through pipe
// The following example creates a Go program, which reads data passed from standard intput via the pipe.
func main() {

	nBytes, nChunks := int64(0), int64(0)
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)

	for {

		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		if n == 0 {

			if err == nil {
				continue
			}

			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		nChunks++
		nBytes += int64(len(buf))

		fmt.Println(string(buf))

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}

	fmt.Println("Bytes:", nBytes, "Chunks:", nChunks)
}

/*
date | go run read_stdin.go
Thu Dec  2 16:00:57 CST 2021

Bytes: 29 Chunks: 1
*/
