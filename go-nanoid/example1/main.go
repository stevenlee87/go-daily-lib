package main

import (
    "fmt"
    "github.com/matoous/go-nanoid/v2"
)

func main() {
    id, err := gonanoid.New()
    if err != nil {
        fmt.Println("Error generating ID:", err)
        return
    }

    fmt.Println("Generated ID:", id)
}
