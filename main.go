package main

import (
	"flag"
	"fmt"

	"github.com/m4nsur/weather-lrn/geo"
	"github.com/m4nsur/weather-lrn/weather"
)

func main() {
	city := flag.String("city", "", "Введите название города")
	format := flag.Int("format", 1, "Формат вывода погоды")
	flag.Parse()
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		println(err)
	}
	data, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}