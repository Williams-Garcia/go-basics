package panic_example

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

var (
	errFailedRead       = &customError{"Error: el archivo indicado no fue encontrado o está dañado"}
	errFailedContent    = &customError{"Error: el archivo indicado no contiene informacion"}
	errClientExist      = &customError{"Error: el cliente ya existe"}
	errClientName       = &customError{"Error: el Nombre esta vacio"}
	errClientLegajo     = &customError{"Error: el Legajo esta vacio"}
	errClientDni        = &customError{"Error: el Dni esta vacio"}
	errClientCellNumber = &customError{"Error: el numero celular esta vacio"}
	errClientAddress    = &customError{"Error: la direccion esta vacia"}
)

type Client struct {
	Legajo     string `json:"legajo"`
	Name       string `json:"name"`
	Dni        string `json:"dni"`
	CellNumber string `json:"cellNumber"`
	Address    string `json:"address"`
}

var generalClient = Client{
	Legajo:     "C",
	Name:       "Cx",
	Dni:        "Cxc",
	CellNumber: "Cvx",
	Address:    "sSQW",
}

type customError struct {
	msg string
}

func (c *customError) Error() string {
	return c.msg
}

func Init() {
	openFile("panic_example/customers_input.txt")
}

func openFile(fileName string) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Ejecucion finalizada")
	}()

	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Errorf("%w", errFailedRead))
	}
	defer file.Close()
	readContent(fileName)
}

func readContent(fileName string) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	dataFile, err := os.ReadFile(fileName)
	if err != nil {
		panic(fmt.Errorf("%w", errFailedRead))
	}

	if string(dataFile) == "" {
		panic(fmt.Errorf("%w", errFailedContent))
	}
	convertData(string(dataFile))
}

func convertData(data string) {
	var dataSlice = make([]Client, 0)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		keyVal := strings.Split(line, ",")
		dataSlice = append(dataSlice, Client{Legajo: keyVal[0], Name: keyVal[1], Dni: keyVal[2], CellNumber: keyVal[3], Address: keyVal[4]})
	}
	for _, client := range dataSlice {
		if validClientExist(client, generalClient) {
			panic(fmt.Errorf("%w", errClientExist))
		}
	}
	addClientToList(generalClient, dataSlice)

}

func validClientExist(client1 Client, client2 Client) bool {
	return reflect.DeepEqual(client1, client2)
}

func validDataClient(client Client) {
	fmt.Println(client.Name)
	if client.Name == "" {
		panic(errClientName)
	}
	if client.Legajo == "" {
		panic(errClientLegajo)
	}
	if client.Dni == "" {
		panic(errClientDni)
	}
	if client.Address == "" {
		panic(errClientAddress)
	}
	if client.CellNumber == "" {
		panic(errClientCellNumber)
	}
}

func addClientToList(newClient Client, listClients []Client) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	validDataClient(newClient)
	listClients = append(listClients, newClient)
	fileOutput, err := json.Marshal(listClients)
	if err != nil {
		panic(err)
	}
	var file, err2 = os.OpenFile("panic_example/customers_output.txt", os.O_RDWR, 0644)
	if err2 != nil {
		panic(err2)
	}
	_, err = file.WriteString(string(fileOutput) + "\n")
	err = file.Sync()
	if err != nil {
		panic(err)
	}
	fmt.Println("Archivo actualizado existosamente.")
}
