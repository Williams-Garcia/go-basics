package bucle

import "fmt"

func bucle() {
	months := []string{"ene", "feb", "jun", "aug", "apr"}
	// getSeasonByIf(months)
	getSeasonBySwitch(months)
}

func getSeasonByIf(months []string) {
	for _, month := range months {
		if month == "ene" || month == "feb" {
			fmt.Println("Summer")
		} else if month == "apr" || month == "jun" {
			fmt.Println("Fall")
		} else if month == "aug" {
			fmt.Println("winter")
		} else {
			fmt.Println("Month no exist")
		}
	}
}

func getSeasonBySwitch(months []string) {
	for _, month := range months {
		switch month {
		case "ene", "feb":
			fmt.Println("Summer")
		case "apr", "jun":
			fmt.Println("Fall")
		case "aug":
			fmt.Println("winter")
		default:
			fmt.Println("Month no exist")
		}
	}
}
