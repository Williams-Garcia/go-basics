package clients

import "fmt"

type Client struct {
	name string
	age int
	old float64
	job bool
	pay int
}

func GenerateClients()  {
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