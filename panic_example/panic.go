package panic_example

import (
	"fmt"
	"log"
	"os"
)

var (
	errFailedRead = &customError{"el archivo indicado no fue encontrado o está dañado"}
)

type customError struct {
	msg string
}

func (c *customError) Error() string {
	return c.msg
}

func Init() {
	readFile("./customers.txt")
}

func readFile(fileName string) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Ejecucion finalizada")
	}()

	_, err := os.ReadFile(fileName)
	if err != nil {
		panic(fmt.Errorf("%w", errFailedRead))
	}
}
