package error

import (
	"errors"
	"fmt"
)

func Init() {
	ok, err := sayHi("Main")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(ok)
}

func sayHi(name string) (string, error) {
	if name == "" {
		return "", errors.New("No name provided")
	}
	return fmt.Sprintf("Hi, %s .", name), nil
}
