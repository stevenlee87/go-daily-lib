package main

import (
	"fmt"
	"time"

	stream2 "github.com/sourcegraph/conc/stream"
)

func ExampleStream() {
	times := []int{20, 52, 16, 45, 4, 80}

	stream := stream2.New()
	for _, millis := range times {
		dur := time.Duration(millis) * time.Millisecond
		stream.Go(func() stream2.Callback {
			time.Sleep(dur)
			// This will print in the order the tasks were submitted
			return func() { fmt.Println(dur) }
		})
	}
	stream.Wait()

	// Output:
	// 20ms
	// 52ms
	// 16ms
	// 45ms
	// 4ms
	// 80ms
}

func main() {
	ExampleStream()
}
