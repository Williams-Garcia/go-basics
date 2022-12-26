package maps

import "fmt"

func maps() {
	myMap := map[string]int{"Benji": 20, "Nahu": 26}
	fmt.Println(myMap)
	myMap["Messi"] = 31
	fmt.Println(myMap)

	for key, val := range myMap {
		fmt.Printf("key: %s, value: %d\n", key, val)
		// fmt.Println(key)
		// fmt.Println(val)
	}
}
