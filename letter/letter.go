package letter

import "fmt"

func Letter()  {
	var word string
	fmt.Println("Ingrese palabra")
	fmt.Scanln(&word)
	for _, letter := range word {
		fmt.Printf("letter: %c\n",letter)
	}
	fmt.Println("longitud", len(word))
}
