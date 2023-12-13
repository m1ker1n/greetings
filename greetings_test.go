package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Igor"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) {
		t.Fatalf(`Hello(%q) = %q, %v; want match for %#q, nil`, name, msg, err, want)
	}
}

func TestEmptyHello(t *testing.T) {
	name := ""
	msg, err := Hello(name)
	if msg != "" || err == nil {
		t.Fatalf(`Hello(%q) = %q, %v; want "", error`, name, msg, err)
	}
}
