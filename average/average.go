package average

import "fmt"

func CalculateAverage(notes ...float64) float64 {
	var sumNotes float64
	for _, note := range notes {
		if note < 0 {
			fmt.Println("No puede haber una nota negativa")
			break
		}
		sumNotes += note
	}
	fmt.Printf("El promedio es: %v\n", sumNotes/float64(len(notes)))
	return sumNotes / float64(len(notes))
}
