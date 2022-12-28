package error_taxes_three

import (
	"errors"
	"fmt"
)

var (
	errSalary = errors.New("Error: el salario es menor a 10000")
)

func Init() {
	salary := 10000
	result, err := valid(salary)
	if err != nil {
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
