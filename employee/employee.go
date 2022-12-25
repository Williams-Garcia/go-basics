package employee

import "fmt"

func employees()  {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	listEmployee(employees)
	employees["Federico"] = 25
	fmt.Println("Agregamos a fede")	
	listEmployee(employees)
	fmt.Println("Borramos a pedro")
	delete(employees, "Pedro")
	listEmployee(employees)
}

func listEmployee(employees map[string]int)  {
	var more21 int = 0
	for key, val := range employees {
		fmt.Printf("employee: %s, age: %d\n", key, val)
		if val >= 21 {
			more21++
		}
	}
	fmt.Printf("Personas con mas de 21 %d\n", more21)
}