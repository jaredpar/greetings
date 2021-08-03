package greetings

import (
	"fmt"
)

func Random() string {
	return "there is no random"
}

func Hello(name string) string {
	message := fmt.Sprintf("Hi %v!", name)
	return message
}
