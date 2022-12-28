package error_taxes_five

import (
	"errors"
	"fmt"
	"log"
)

var (
	errSalary = errors.New("Error: el mínimo imponible es de 150.000 y el salario ingresado es de:")
	errHours  = errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
)

func Init() {
	salary := 155000.0
	hours := 20.0
	/* Manejando el error y validando con panic en main
	result, err := calculateSalaryPerMonthByHours(salary, hours)
	if err != nil {
		if errors.Is(err, errSalary) {
			panic(err)
		}
		if errors.Is(err, errHours) {
			panic(err)
		}
	}
	fmt.Printf("Salario inicial: %v\n Salario final: %v\n", salary, result)*/

	//Por manejo de panic con defer y recover
	result := calculateSalaryPerMonthByHoursPanic(salary, hours)
	if result != 0 {
		fmt.Printf("Salario inicial: %v\n Salario final: %v\n", salary, result)
	}
}

func calculateSalaryPerMonthByHours(salary float64, workHours float64) (totalSalary float64, err error) {
	if salary <= 150000.0 {
		return 0, fmt.Errorf("%w %v", errSalary, salary)
	}

	if workHours < 80.0 {
		return 0, fmt.Errorf("%w", errHours)
	}

	return salary - (salary * 0.10), nil
}

func calculateSalaryPerMonthByHoursPanic(salary float64, workHours float64) (totalSalary float64) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()

	if salary <= 150000.0 {
		panic(fmt.Errorf("%w %v", errSalary, salary))
	}

	if workHours < 80.0 {
		panic(fmt.Errorf("%w", errHours))
	}

	return salary - (salary * 0.10)
}
