package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf(randomFormat(), name), nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %s. Welcome!",
		"Great to see you, %s!",
		"Hail, %s! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
