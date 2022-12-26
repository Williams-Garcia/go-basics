package weather

import "fmt"

var temperature float64 = 24.8
var humidity float64 = 5.0
var pressure float64 = 3.12

func PrintWeather() {
	fmt.Println(temperature)
	fmt.Println(humidity)
	fmt.Println(pressure)
}
