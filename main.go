package main

import (
	"fmt"

	"github.com/wneessen/go-mail"
)

func PartToString(p *mail.Part) string {
	body, _ := p.GetContent()
	return string(body)
}

type Config struct {
	sender    string
	recipient string
	subject   string
	body      string
}

func GetDefaultConfig() Config {
	return Config{
		sender:    "test@example.com",
		recipient: "test@example.com",
		subject:   "Test",
		body:      "Hello World!",
	}
}

type ConfigLambda func(*Config)

func WithSubject(s string) ConfigLambda {
	return func(c *Config) {
		c.subject = s
	}
}

func WithSender(s string) ConfigLambda {
	return func(c *Config) {
		c.sender = s
	}
}

func WithRecipient(s string) ConfigLambda {
	return func(c *Config) {
		c.recipient = s
	}
}

func WithBody(s string) ConfigLambda {
	return func(c *Config) {
		c.body = s
	}
}

func CreateMessage(cls ...ConfigLambda) *mail.Msg {
	message := mail.NewMsg()
	config := GetDefaultConfig()
	for _, lambda := range cls {
		lambda(&config)
	}
	message.SetBodyString(mail.TypeTextPlain, config.body)
	message.Subject(config.subject)
	message.From(config.sender)
	message.To(config.recipient)
	return message
}

func main() {
	message := CreateMessage()
	body := message.GetParts()[0]
	fmt.Println(PartToString(body))
}
