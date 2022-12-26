package structures

import (
	"fmt"
	// "encoding/json"
)

type Person struct {
	name string
	age   uint
	position Position
}

// type Person struct {
// 	name string `json:"name"`
// 	age   int `json:"age"`
// 	position Position `json:"position"`
// }

type Position struct {
	position string
	number uint
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