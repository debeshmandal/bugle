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
