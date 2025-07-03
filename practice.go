package something

import (
	"errors"
	"fmt"
)

func Something(name string) (string, error) {
	message := fmt.Sprintf("Hi, %v. Welcome to Chilli's!", name)
	if message == "" {
		return "", errors.New("you forgot the name dumbass")

	}

	return message, nil
}
