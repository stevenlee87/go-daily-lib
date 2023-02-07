package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/sourcegraph/conc/pool"
)

func ExampleContextPool_WithCancelOnError() {
	p := pool.New().
		WithMaxGoroutines(4).
		WithContext(context.Background()).
		WithCancelOnError()
	for i := 0; i < 3; i++ {
		i := i
		p.Go(func(ctx context.Context) error {
			fmt.Println(i)
			if i == 2 {
				return errors.New("I will cancel all other tasks!")
			}
			<-ctx.Done()
			return nil
		})
	}
	err := p.Wait()
	fmt.Println(err)
	// Output:
	// I will cancel all other tasks!
}

func main() {
	ExampleContextPool_WithCancelOnError()
}
