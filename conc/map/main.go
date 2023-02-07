package main

import (
	"fmt"
	"github.com/sourcegraph/conc/iter"
)

func main() {
	input := []int{1, 2, 3, 4}
	mapper := iter.Mapper[int, bool]{
		MaxGoroutines: len(input) / 2,
	}

	results := mapper.Map(input, func(v *int) bool { return *v%2 == 0 })
	fmt.Println(results)
	// Output:
	// [false true false true]
}
