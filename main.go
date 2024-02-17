package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
	if err := message.From(config.sender); err != nil {
		log.Fatalf("Failed to set From address: %s", err)
	}
	if err := message.To(config.recipient); err != nil {
		log.Fatalf("Failed to set To address: %s", err)
	}
	return message
}

func CreateClient() *mail.Client {
	smtp_address := os.Getenv("BUGLE_SMTP_SERVER")
	client, err := mail.NewClient(
		smtp_address,
		mail.WithPort(587),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(os.Getenv("BUGLE_USERNAME")),
		mail.WithPassword(os.Getenv("BUGLE_PASSWORD")),
	)
	if err != nil {
		log.Fatalf("Failed to create mail client: %s", err)
	}
	return client
}

func main() {
	var body string
	var sender string
	var recipient string
	var subject string
	var dryRun bool

	flag.StringVar(&body, "body", "", "E-mail Body")
	flag.StringVar(&subject, "subject", "", "E-mail Subject")
	flag.StringVar(&sender, "sender", "", "E-mail Sender")
	flag.StringVar(&recipient, "recipient", "", "E-mail Recipient")
	flag.BoolVar(&dryRun, "dry-run", false, "Set to skip using client, useful for testing and CI")

	flag.Parse()

	message := CreateMessage(
		WithSender(sender),
		WithSubject(subject),
		WithRecipient(recipient),
		WithBody(body),
	)

	if dryRun {
		fmt.Println(PartToString(message.GetParts()[0]))
	} else {
		client := CreateClient()
		if err := client.DialAndSend(message); err != nil {
			log.Fatalf("Failed to send mail: %s", err)
		}
	}
}
