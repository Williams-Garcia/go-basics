package students

import "fmt"

type Student struct {
	dni        uint
	name       string
	lastName   string
	dateOfJoin string
}

func (s Student) printDataStudent() {
	fmt.Println(s)
}

func Init() {
	s := Student{
		dni:        1120,
		name:       "Juan",
		lastName:   "Perez",
		dateOfJoin: "22-12-2020",
	}
	s.printDataStudent()
}
