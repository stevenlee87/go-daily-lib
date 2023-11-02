package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// 创建一个 map 用于存储元素的值与它们的索引的映射
	hashTable := make(map[int]int)

	for i, x := range nums {
		// 计算目标值与当前元素的差值
		complement := target - x

		// 检查 map 中是否存在这个差值
		if p, ok := hashTable[complement]; ok {
			fmt.Println("p is:", p)
			// 如果存在，返回差值的索引和当前元素的索引
			return []int{p, i}
		}

		// 否则，将当前元素的值作为键，索引作为值存入 map
		hashTable[x] = i
		fmt.Printf("hashTable[%s] is: %d\n", fmt.Sprint(x), hashTable[x])
	}
	return nil
}

func main() {
	//nums := []int{2, 7, 11, 15}
	//target := 9
	nums := []int{2, 7, 11, 15}
	target := 13
	result := twoSum(nums, target)
	if result != nil {
		fmt.Println("Indices of the two numbers:", result[0], result[1])
	} else {
		fmt.Println("No solution found.")
	}
}
