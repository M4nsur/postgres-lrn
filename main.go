package main

import (
	"flag"
	"fmt"

	"github.com/m4nsur/weather-lrn/geo"
)

func main() {


	city := flag.String("city", "", "Введите название города")
	// format := flag.Int("format", 1, "Формат вывода погоды")
	flag.Parse()
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		println(err)
	}
	fmt.Println(geoData.City)
}