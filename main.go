package main

import (
	"fmt"
	"live_class/animals"
	"live_class/average"
	"live_class/clients"
	"live_class/ecommerce"
	"live_class/fullName"
	"live_class/letter"
	"live_class/month"
	"live_class/numbers"
	"live_class/salary"
	"live_class/statistics"
	"live_class/taxes"
	"live_class/weather"
) //Logs

// var number = 10

// const otherNumber = 20

func main() {
	fmt.Println(menu())
	var eleccion int //Declarar variable y tipo antes de escanear, esto es obligatorio
	fmt.Scanln(&eleccion)
	switchOption(eleccion)
	// var (
	// 	a,b,c float64
	// 	name, lastName string
	// )
	//declaracion corta
	// if var eleccion; condicion {caso true}
	//arrays
	//var x [y]type / x[0] = y
	//slices
	//var s = []type{y,y1}
	//a:= make([]type, size)
	//[1:n]
}

func menu() string {
	menu :=
		"Hola!\nQue clase de git te gustaria ejecutar?\n" +
			"[1] Hola Mundo\n[2] Variables\n[3] GO dia 1\n" +
			"[4] Go dia 2\n[5] Go dia 3\n"

	return menu
}

func ifOption(eleccion int) {
	if eleccion == 1 {
		fmt.Println("Hola Mundo")
		myPi := numbers.Pi
		fmt.Println(myPi)
	} else if eleccion == 2 {
		fullName.PrintFullName()
		fullName.PrintAddress()
		weather.PrintWeather()
	} else {
		fmt.Println("Clase aun no disponible")
	}
}

func switchOption(eleccion int) {
	switch eleccion {
	case 1:
		fmt.Println("Hola Mundo")
		myPi := numbers.Pi
		fmt.Println(myPi)
	case 2:
		fullName.PrintFullName()
		fullName.PrintAddress()
		weather.PrintWeather()
	case 3:
		letter.Letter()
		clients.GenerateClients()
		month.GetMonthByNumber()
	case 4:
		taxes.PrintMenu()
		average.CalculateAverage(1, 2)
		salary.IdentifySalary("c", 150)
		statistics.IdentifyOperation("average", 2, 4, 1)
		animals.Init()
	case 5:
		// structures.Init()
		// products.Init()
		// manage.Init()
		// example_pointers.Init()
		// example_interfaces.Init()
		// empty_interfaces.Init()
		// students.Init()
		ecommerce.Init()
	default:
		fmt.Println("Clase aun no disponible")
	}
}
