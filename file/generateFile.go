package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Set a seed for the random number generator
	rand.Seed(time.Now().UnixNano())

	// Open the file for writing
	file, err := os.Create("file/largefile.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Generate 10,000 lines of random content
	for i := 0; i < 100; i++ {
		line := generateRandomContent() + "\n"
		file.WriteString(line)
	}
	fmt.Println("File generation complete.")
}

// Function to generate random content for each line
func generateRandomContent() string {
	// Define characters for random content
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Generate a random string of length 50
	randomContent := make([]byte, 50)
	for i := range randomContent {
		// fmt.Println("i is:", i)
		randomContent[i] = chars[rand.Intn(len(chars))]
		// fmt.Println(rand.Intn(len(chars)))
	}
	return string(randomContent)
}
