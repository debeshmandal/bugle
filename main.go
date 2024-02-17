package main

import (
	"fmt"

	"github.com/wneessen/go-mail"
)

func PartToString(p *mail.Part) string {
	body, _ := p.GetContent()
	return string(body)
}

func CreateMessage() *mail.Msg {
	message := mail.NewMsg()
	message.SetBodyString(mail.TypeTextPlain, "Hello World!")
	return message
}

func main() {
	message := CreateMessage()
	body := message.GetParts()[0]
	fmt.Println(PartToString(body))
}
