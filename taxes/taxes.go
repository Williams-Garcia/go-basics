package taxes

import "fmt"

var taxesSalary float64 = 0

func PrintMenu() {
	fmt.Println("Ingrese salario para calcular impuesto")
	var salary float64
	fmt.Scanln(&salary)
	fmt.Printf("El impuesto del salario: %v\n es: %v\n", salary, calculateTaxes(salary))
}

func calculateTaxes(salary float64) float64 {
	if salary > 50000 {
		taxesSalary = salary * .17
	} else if salary > 150000 {
		taxesSalary = salary * .27
	}
	return taxesSalary
}
