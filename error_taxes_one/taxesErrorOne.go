package error_taxes_one

import (
	"errors"
	"fmt"
)

var (
	errSalary = errors.New("Error: el salario ingresado no alcanza el mínimo imponible")
)

func Init() {
	salary := 140000
	result, err := valid(salary)
	if err != nil {
		fmt.Println(errors.Is(err, errSalary))
		// fmt.Errorf("%w, %s", err, ErrSalary)
		panic(err)
	}
	fmt.Println(result)

}

func valid(salary int) (msg string, err error) {
	if salary < 150000 {
		return "", errSalary
	}
	return "Debe pagar impuesto", nil
}
