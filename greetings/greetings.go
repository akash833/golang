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
	message := fmt.Sprintf(randomPrint(), name)
	return message, nil
}

func randomPrint() string {
	formats := []string{
		"Hii, %v.Welcome",
		"Hello, %v. How are you?",
		"Hyy, %v",
	}
	return formats[rand.Intn(len(formats))]
}
