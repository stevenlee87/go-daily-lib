package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	/*
		// WithSeconds overrides the parser used for interpreting job schedules to
		// include a seconds field as the first one.
		func WithSeconds() Option {
			return WithParser(NewParser(
				Second | Minute | Hour | Dom | Month | Dow | Descriptor,
			))
		}
	*/

	c := cron.New(cron.WithSeconds())
	c.AddFunc("1 * * * * *", func() {
		fmt.Println("every 1 second")
	})
	c.Start()

	time.Sleep(62 * time.Second)
}
