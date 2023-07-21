package main

import "fmt"

func findMaxAndSecondMax(nums []int) (int, int) {
	if len(nums) < 2 {
		return 0, 0
	}

	max := nums[0]
	secondMax := nums[1]

	if secondMax > max {
		max, secondMax = secondMax, max
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > max {
			secondMax = max
			max = nums[i]
		} else if nums[i] > secondMax {
			secondMax = nums[i]
		}
	}

	return max, secondMax
}

func main() {
	nums := []int{5, 10, 3, 8, 15, 7, 25, 12, 6, 18, 30, 9, 20, 4, 22, 14, 28, 16, 2, 19}
	fmt.Println(findMaxAndSecondMax(nums))
}
