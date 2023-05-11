package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name, rand.Intn(3))
	return message, nil
}

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string{
	formats := []string{
		"Hi, %v. Welcome! %t",
		"Great to see you, %v! %t",
		"Hail, %v! Weel met! %t",
	}

	return formats[rand.Intn(len(formats))]
}