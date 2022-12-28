package error_taxes_four

import (
	"errors"
	"fmt"
)

var (
	errSalary = errors.New("Error: el mínimo imponible es de 150.000 y el salario ingresado es de:")
)

func Init() {
	salary := 150000
	result, err := valid(salary)
	if err != nil {
		if errors.Is(err, errSalary) {
			panic(err)
		}
	}
	fmt.Println(result)

}

func valid(salary int) (msg string, err error) {
	if salary <= 150000 {
		return "", fmt.Errorf("%w %d", errSalary, salary)
	}
	return "Debe pagar impuesto", nil
}
