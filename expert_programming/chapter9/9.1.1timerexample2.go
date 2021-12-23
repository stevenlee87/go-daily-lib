package main

import (
	"log"
	"time"
)

func DelayFunction() {
	timer := time.NewTimer(5 * time.Second)

	select {
	case <-timer.C:
		log.Println("Delayed 5s, start to do something.")
	}
}

func main() {
	DelayFunction()
}
