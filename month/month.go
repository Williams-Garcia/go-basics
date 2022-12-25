package month

import "fmt"

func GetMonthByNumber()  {
	var monthNumber int
	fmt.Println("Ingrese mes")
	fmt.Scanln(&monthNumber)
	
	switch monthNumber {
		case 1:
			fmt.Printf("month: %s, number: %d\n", "ene", monthNumber)
		case 2:
			fmt.Printf("month: %s, number: %d\n", "feb", monthNumber)
		case 3:
			fmt.Printf("month: %s, number: %d\n", "mar", monthNumber)
		case 4:
			fmt.Printf("month: %s, number: %d\n", "apr", monthNumber)
		case 5:
			fmt.Printf("month: %s, number: %d\n", "may", monthNumber)
		case 6:
			fmt.Printf("month: %s, number: %d\n", "jun", monthNumber)
		case 7:
			fmt.Printf("month: %s, number: %d\n", "jul", monthNumber)
		case 8:
			fmt.Printf("month: %s, number: %d\n", "aug", monthNumber)
		case 9:
			fmt.Printf("month: %s, number: %d\n", "sep", monthNumber)
		case 10:
			fmt.Printf("month: %s, number: %d\n", "oct", monthNumber)
		case 11:
			fmt.Printf("month: %s, number: %d\n", "nov", monthNumber)
		case 12:
			fmt.Printf("month: %s, number: %d\n", "dec", monthNumber)
		default:
			fmt.Println("Month no exist")
	}
}