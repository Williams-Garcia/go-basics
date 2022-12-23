package main

import (
	"fmt"
	"live_class/fullName"
	"live_class/weather"
	"live_class/numbers"
	// "live_class/animals"
	// "live_class/statistics"
	// "live_class/taxes"
	// "live_class/average"
	"live_class/salary"
) //Logs

// var number = 10

// const otherNumber = 20

func main() {
	// animals.Init()
	// statistics.IdentifyOperation("average", 2, 4, 1)
	// minValue := statistics.minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	// averageValue := statistics.averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	// maxValue := statistics.maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	// fmt.Printf("El minimo es: %f", minValue)
	// fmt.Printf("El maximo es: %f", maxValue)
	// fmt.Printf("El promedio es: %f", averageValue)
	// bucle()
	// fmt.Print(menu())
	// letter()
	// generateClients()
	// getMonthByNumber()
	// var eleccion int //Declarar variable y tipo antes de escanear, esto es obligatorio
	// fmt.Scanln(&eleccion)
	// // ifOption(eleccion)
	// switchOption(eleccion)
	// var (
	// 	a,b,c float64
	// 	name, lastName string
	// )
	// employees()
	//declaracion corta
	// if var eleccion; condicion {caso true}
	// taxes.PrintMenu()
	// average.CalculateAverage(1,2,)
	salary.IdentifySalary("c",150)
	//arrays
	//var x [y]type / x[0] = y 
	//slices
	//var s = []type{y,y1}
	//a:= make([]type, size)
	//[1:n]
	// slices()
	// maps()
}

func letter()  {
	var word string
	fmt.Scanln(&word)
	for _, letter := range word {
		fmt.Printf("letter: %c\n",letter)
	}
	fmt.Println("longitud", len(word))
}

type Client struct {
	name string
	age int
	old float64
	job bool
	pay int
}

func generateClients()  {
	clients := []Client{}

	clientOne := Client{
		"pp", 22, 1, false, 100,
	}
	clientTwo := Client{
		"jj", 18, 2, true, 90,
	}
	clientThree := Client{
		"bv", 35, 10, true, 200,
	}

	clients = append(clients, clientOne)
	clients = append(clients, clientTwo)
	clients = append(clients, clientThree)

	// fmt.Println(clients)
	for _, client := range clients {
		validateCases(client)
	}
}

func validateCases(client Client)  {
	if client.age >= 22 && client.old >= 1 && client.job == true{
		if client.pay >= 100 {
			fmt.Printf("c% pago sin interes \n", client.name)
		}else{
			fmt.Printf("c% pago con inteeres \n",  client.name)
		}
	} else if client.age < 22 {
		fmt.Printf("c%  No cumples con la edad \n", client.name)
	} else if client.job != true {
		fmt.Printf("c% No cumples con un trabajo \n", client.name)
	} else {
		fmt.Printf("c% no eres apto para credito \n", client.name)
	}
	
}

func getMonthByNumber()  {
	var monthNumber int
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

func slices(){
	slc := []int{1,2,3,4,5}
	fmt.Println( slc)
}

func maps()  {
	myMap := map[string]int{"Benji" : 20, "Nahu" : 26}
	fmt.Println(myMap)
	myMap["Messi"] = 31
	fmt.Println(myMap)

	for key, val := range myMap {
		fmt.Printf("key: %s, value: %d\n", key, val)
		// fmt.Println(key)
		// fmt.Println(val)
	}
}

func menu() string {
	menu := 
	"Hola!\nQue clase de git te gustaria ejecutar?\n[1] Hola Mundo\n[2] Variables\n"
	
	return menu
}

func ifOption(eleccion int)  {
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

func switchOption(eleccion int)  {
	switch eleccion {
	case 1:
		fmt.Println("Hola Mundo")
		myPi := numbers.Pi
		fmt.Println(myPi)

	case 2:
		fullName.PrintFullName()
		fullName.PrintAddress()
		weather.PrintWeather()
	
	default:
		fmt.Println("Clase aun no disponible")
	}
}

func bucle()  {
	months := []string{"ene","feb","jun","aug","apr"}
	// getSeasonByIf(months)
	getSeasonBySwitch(months)
}

func getSeasonByIf(months []string)  {
	for _, month := range months{
		if month == "ene" || month == "feb" {
			fmt.Println("Summer")
		} else if month == "apr" || month == "jun" {
			fmt.Println("Fall")
		} else if month == "aug" {
			fmt.Println("winter")
		} else{
			fmt.Println("Month no exist")
		}
	}
}

func getSeasonBySwitch(months []string)  {
	for _, month := range months{
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