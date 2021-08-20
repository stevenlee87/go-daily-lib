package sort

func InsertionSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 1; i < length; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j--
		}
	}
	return arr
}
