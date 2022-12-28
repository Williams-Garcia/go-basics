package error

import (
	"errors"
	"fmt"
)

type customError struct {
	msg  string
	code int
}

func (c customError) Error() string {
	return c.msg
}

func Init3() {
	err := &customError{
		msg:  "failed",
		code: 500,
	}

	fmt.Println("error: ", err)

	var ce *customError
	fmt.Println(ce)
	fmt.Println(errors.As(err, &ce))

	errAnid := fmt.Errorf("%w, extra details", err)
	fmt.Println("anid: ", errors.As(errAnid, &ce))
}
