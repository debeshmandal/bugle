package main

import (
	"regexp"
	"testing"
)

func TestCreateMessage(t *testing.T) {
	test_body := regexp.MustCompile("Hello World!")
	msg := CreateMessage()
	body := PartToString(msg.GetParts()[0])
	if !test_body.MatchString(body) {
		t.Fatalf(`CreateMessage() = %q; want match for %#q`, body, test_body)
	}
}

func TestWithLambda(t *testing.T) {
	test_config := Config{
		body:      "Hello",
		sender:    "Hello@example.com",
		recipient: "Hello@example.com",
		subject:   "Hello",
	}
	msg := CreateMessage(
		WithBody(test_config.body),
		WithSender(test_config.sender),
		WithSubject(test_config.subject),
		WithRecipient(test_config.recipient),
	)
	body := PartToString(msg.GetParts()[0])
	if !regexp.MustCompile(test_config.body).MatchString(body) {
		t.Fatalf(`CreateMessage() = %q; want match for %#q`, body, test_config.body)
	}
}
