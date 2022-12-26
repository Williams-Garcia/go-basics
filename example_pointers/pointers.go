package example_pointers

import "fmt"

func increase(v *int) {
	*v++
}

func Init() {
	var v int = 19
	increase(&v)
	fmt.Println("El valor de v ahora vale: ", v)
}
