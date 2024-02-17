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
	test_config := GetDefaultConfig()
	msg := CreateMessage(
		WithBody(test_config.body),
		WithSender(test_config.body),
		WithSubject(test_config.body),
		WithRecipient(test_config.body),
	)
	body := PartToString(msg.GetParts()[0])
	if !regexp.MustCompile(test_config.body).MatchString(body) {
		t.Fatalf(`CreateMessage() = %q; want match for %#q`, body, test_config.body)
	}
}
