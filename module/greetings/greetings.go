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

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func RandomGreeting(name string) string {

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	message := fmt.Sprintf(formats[rand.Intn(len(formats))], name)
	return message
}
