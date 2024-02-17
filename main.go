package main

import (
	"flag"
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
		if s != "" {
			c.subject = s
		}
	}
}

func WithSender(s string) ConfigLambda {
	return func(c *Config) {
		if s != "" {
			c.sender = s
		}
	}
}

func WithRecipient(s string) ConfigLambda {
	return func(c *Config) {
		if s != "" {
			c.recipient = s
		}
	}
}

func WithBody(s string) ConfigLambda {
	return func(c *Config) {
		if s != "" {
			c.body = s
		}
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
	var body string
	var sender string
	var recipient string
	var subject string

	flag.StringVar(&body, "body", "", "E-mail Body")
	flag.StringVar(&subject, "subject", "", "E-mail Subject")
	flag.StringVar(&sender, "sender", "", "E-mail Sender")
	flag.StringVar(&recipient, "recipient", "", "E-mail Recipient")

	flag.Parse()

	message := CreateMessage(
		WithBody(body),
		WithSender(sender),
		WithSubject(subject),
		WithRecipient(recipient),
	)

	fmt.Println(PartToString(message.GetParts()[0]))
}
