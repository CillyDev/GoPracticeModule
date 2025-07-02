package something

import "fmt"

func Something(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome to Chilli's!", name)
	fmt.Println(message)
	return (message)
}
