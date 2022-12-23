package salary

import "fmt"

var salaryPerHour float64

func IdentifySalary(category string, hoursInMinutes int) float64 {
	var hours float64 = float64(hoursInMinutes)/60
	if hours > 1{
		switch category {
		case "A", "a":
			salaryPerHour = (3000 * hours) + ((3000 * 160)* .50)
		case "B", "b":
			salaryPerHour = (1500 * hours) + ((1500 * 160)* .20)
		case "C", "c":
			salaryPerHour = 1000 * hours
		default:
			fmt.Println("No existe dicha categoria")
		}
		fmt.Printf("Categoria: %s, Salario: %v\n", category, salaryPerHour)
		return salaryPerHour
	}
	fmt.Printf("No has trabajado el tiempo minimo")
	return 0
}
