package empty_interfaces

import "fmt"

type ListHeterogen struct {
	Data []interface{}
}

func assertionType() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)
	f = i.(float64)
	fmt.Println(f)
}

func Init() {
	l := ListHeterogen{}
	l.Data = append(l.Data, 1)
	l.Data = append(l.Data, "Hola")
	l.Data = append(l.Data, true)

	fmt.Println("%v\n", l.Data)

	assertionType()
}
