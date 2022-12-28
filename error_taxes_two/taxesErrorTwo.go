package error_taxes_two

import (
	"errors"
	"fmt"
)

var (
	errSalary = &customError{"Error: el salario es menor a 10000"}
)

type customError struct {
	msg string
}

func (c *customError) Error() string {
	return c.msg
}

func Init() {
	salary := 10000
	result, err := valid(salary) //validamos salario
	if err != nil {              //agregar error
		err = fmt.Errorf("%w", errSalary) //anido el error
		if errors.Is(err, errSalary) {    //despues se compara
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
