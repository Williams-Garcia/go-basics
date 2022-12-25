package structures

import "fmt"

type Person struct {
	name string
	age   int
	position Position
}

type Position struct {
	position string
	number int
}

func Init()  {
	getDescription()
}
func getDescription() {
	p := Person{
		name: "Gabriel",
		age:   22,
		position: Position{"GoalKeeper", 1},
	}

	p.Description()
}

func (p *Person) Description() {
	fmt.Printf("%s tiene %d a~nos de edad y es %s con el numero: %d\n", p.name, p.age, p.position.position, p.position.number)
}