package main

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {
	e := email.NewEmail()
	e.From = "stevenlee87 <stevenlee87@126.com>"
	e.To = []string{"stevenlee87@126.com"}
	e.Subject = "Go 每日一库"
	e.Text = []byte("请看附件")
	e.AttachFile("test.txt")
	err := e.Send("smtp.126.com:25", smtp.PlainAuth("", "stevenlee87@126.com", "123456", "smtp.126.com"))
	if err != nil {
		log.Fatal("failed to send email:", err)
	}
}
