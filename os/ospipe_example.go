package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	inR, inW, _ := os.Pipe()
	outR, outW, _ := os.Pipe()
	done := make(chan struct{})

	process, _ := os.StartProcess("/bin/sh", nil, &os.ProcAttr{
		Files: []*os.File{inR, outW, outW}})

	go func() {
		writer := bufio.NewWriter(inW)
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			writer.WriteString("date\n")
			writer.Flush()
		}
		inW.Close()
		outW.Close()
	}()
	go func() {
		scanner := bufio.NewScanner(outR)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		process.Signal(os.Kill)
		done <- struct{}{}
		fmt.Println("finish")
	}()

	process.Wait()
	<-done

}
