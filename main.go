package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("хелло")

	city := flag.String("city", "", "Введите название города")
	format := flag.Int("format", 1, "Формат вывода погоды")
	flag.Parse()
	fmt.Println(*city, *format)
}