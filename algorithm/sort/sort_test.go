package sort

import (
	"testing"
)

func testFramework(t *testing.T, sortingFunction func([]int) []int) {
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			actual := sortingFunction(test.input)
			pos, sorted := compareSlices(actual, test.expected)
			if !sorted {
				if pos == -1 {
					t.Errorf("test %s failed due to slice length changing", test.name)
				}
				t.Errorf("test %s failed at index %d", test.name, pos)
			}
		})
	}
}

//BEGIN TESTS

func TestBubble(t *testing.T) {
	testFramework(t, InsertionSort)
}
