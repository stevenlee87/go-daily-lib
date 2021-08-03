package main

func Utf8Count(s string) int {
	var count int
	for range s {
		count++
	}
	return count
}

fmt.Println(Utf8Count("中国")) // 2

