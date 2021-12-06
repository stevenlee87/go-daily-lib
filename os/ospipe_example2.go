package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	inR, inW, _ := os.Pipe()
	outR, outW, _ := os.Pipe()
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	process, _ := os.StartProcess("/bin/sh", nil, &os.ProcAttr{
		Files: []*os.File{inR, outW, outW},
		Dir:   dir,
	})

	go func() {
		writer := bufio.NewWriter(inW)
		writer.WriteString("ls -a ~/project/golang/go-daily-lib/os")
		writer.Flush()
		inW.Close()
		outW.Close()
	}()

	process.Wait()
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(outR)
	fmt.Println(buffer.String())
}
