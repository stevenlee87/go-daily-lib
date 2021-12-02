package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Go pipe simple example
// The following example demonstrates the usage of the io.Pipe function.
func main() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "Hello there\n")
		w.Close()
	}()

	_, err := io.Copy(os.Stdout, r)

	if err != nil {
		log.Fatal(err)
	}
}

/*
In the example, we create a pipe with io.Pipe. We write to the pipe's writer in a goroutine and then copy the data from
the pipe's reader to the standard output using io.Copy.

go func() {
    fmt.Fprint(w, "Hello there\n")
    w.Close()
}()
We write some data to the pipe's writer in a goroutine. Each write to the PipeWriter blocks until it has satisfied
one or more reads from the PipeReader that fully consume the written data.

$ go run simple.go
Hello there
This is the output.
*/
