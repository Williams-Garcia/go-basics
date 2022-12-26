package manage

// Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un pequeño programa nos ayudará a gestionar correctamente dichos empleados. Los objetivos son:
// Crear una estructura Person con los campos ID, Name, DateOfBirth.
// Crear una estructura Employee con los campos: ID, Position y una composicion con la estructura Person.
// Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar la impresión de los campos de un empleado.
// Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último ejecutar el método PrintEmployee().
// Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.

import "fmt"

type Person struct {
	id          uint
	name        string
	dateOfBirth string
}

type Employee struct {
	id       uint
	position string
	person   Person
}

func (e Employee) printDataEmployee() {
	fmt.Println(e)
}

func Init() {
	p := Person{
		id:          1,
		name:        "Juan",
		dateOfBirth: "19-09-1999",
	}
	e := Employee{
		id:       10,
		position: "manager",
		person:   p,
	}
	e.printDataEmployee()
}
