package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("file/largefile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	const maxScanTokenSize = 64 * 1024 * 1024 // 64MB
	buf := make([]byte, maxScanTokenSize)

	scanner := bufio.NewScanner(file)
	scanner.Buffer(buf, maxScanTokenSize)

	for scanner.Scan() {
		line := scanner.Text()
		// 处理每一行的逻辑
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
