package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomString generates a random string of length n
func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// randomInt generates a random integer between min and max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// run  execute linux command
func run(binFile, arg string) error {
	cmd := exec.Command(binFile, arg)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	for {
		text := randomString(randomInt(1, 31))
		fmt.Println(text)
		run("/challenge/main", text)
		time.Sleep(10 * time.Millisecond)
	}
}
