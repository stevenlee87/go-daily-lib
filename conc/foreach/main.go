package main

import (
	"fmt"
	"github.com/sourcegraph/conc/iter"
)

func main() {
	input := []int{1, 2, 3, 4}
	iterator := iter.Iterator[int]{
		MaxGoroutines: len(input) / 2,
	}

	iterator.ForEach(input, func(v *int) {
		if *v%2 != 0 {
			*v = -1
		}
	})

	fmt.Println("10%3 = ", 10%3)     // =1
	fmt.Println("-10%3 = ", -10%3)   // = -10 - (-3*3) = -10 - (-9) = -1
	fmt.Println("10%-3 = ", 10%-3)   // =1
	fmt.Println("-10%-3 = ", -10%-3) // =-1
	fmt.Println("7%-3 = ", 7%-3)     // =1
	fmt.Println("-7%3 = ", -7%3)     // =-1

	fmt.Println(input)
}
