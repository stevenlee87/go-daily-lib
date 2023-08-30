package main

import (
	"encoding/json"
	"log"
)

type S struct {
	A int
	B *int
	C float64
	D func() string
	E chan struct{}
}

func main() {
	s := S{
		A: 1,
		B: nil,
		C: 12.15,
		D: func() string {
			return "Hello"
		},
		E: make(chan struct{}),
	}

	_, err := json.Marshal(s)
	if err != nil {
		log.Printf("err occurred.")
		return
	}
	log.Printf("everything is ok.")

	return
}
