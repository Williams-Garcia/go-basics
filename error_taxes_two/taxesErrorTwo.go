package error_taxes_two

import (
	"errors"
	"fmt"
)

var (
	errSalary = errors.New("Error: el salario es menor a 10000")
)

type customError struct {
	msg string
}

func (c *customError) Error() string {
	return c.msg
}

func Init() {
	salary := 10000
	result, err := valid(salary)
	if err != nil {
		err = &customError{err.Error()}   //agregar error
		err = fmt.Errorf("%w", errSalary) //anido el error
		// fmt.Println(fmt.Errorf("%w, %s", err, errSalary))        //despues se compara
		// fmt.Println(errors.Is(err, errSalary))
		if errors.Is(err, errSalary) {
			panic(err)
		}
	}
	fmt.Println(result)

}

func valid(salary int) (msg string, err error) {
	if salary <= 10000 {
		return "", errSalary
	}
	return "Debe pagar impuesto", nil
}
