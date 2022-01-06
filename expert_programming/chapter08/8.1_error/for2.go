package main

func example() bool {
	err := 1
	i := 0
	for true {
		if i == 2 {
			return true
		}

		i = i + 1
	}
	return false
}

func main() {
	example()
}
