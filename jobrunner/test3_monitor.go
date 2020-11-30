package main

import (
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("Hello,", g.Name)
}

type EmailJob struct {
	Email string
}

func (e EmailJob) Run() {
	fmt.Println("Send,", e.Email)
	time.Sleep(3 * time.Second) // 执行的时候会输出running
}

func main() {
	r := gin.Default()

	jobrunner.Start()
	jobrunner.Every(5*time.Second, GreetingJob{Name: "steven"})
	jobrunner.Every(10*time.Second, EmailJob{Email: "stevenlee87@126.com"})

	r.GET("/jobrunner/json", JobJson)
	r.GET("/jobrunner/html", JobHtml)

	r.Run(":8888")
}

func JobJson(c *gin.Context) {
	c.JSON(200, jobrunner.StatusJson())
}

func JobHtml(c *gin.Context) {
	t, err := template.ParseFiles(os.Getenv("GOPATH") + "/go-daily-lib/jobrunner/views/Status.html")
	if err != nil {
		c.JSON(400, "error")
	}
	t.Execute(c.Writer, jobrunner.StatusPage())
}
