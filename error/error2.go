package error

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("Error: item was not found")
	ErrUnknow   = errors.New("general error")
)

//erros.Is es para errores anidados

func Init2() {
	err := SendError(true)
	errDetail := fmt.Errorf("%w, %s", err, "here are the details")
	fmt.Println(errors.Is(err, ErrNotFound))
	fmt.Println(errors.Is(errDetail, ErrNotFound))
}

func SendError(b bool) error {
	if b {
		return ErrNotFound
	}
	return nil
}
