package main

import (
	"regexp"
	"testing"
)

func TestCreateMessage(t *testing.T) {
	test_message := regexp.MustCompile("Hello World!")
	msg, err := CreateMessage()
	if !test_message.MatchString(msg) || err != nil {
		t.Fatalf(
			`Hello("Gladys") = %q, %v, want match for %#q, nil`,
			msg, err, test_message)
	}
}
