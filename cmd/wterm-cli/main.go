package main

import (
	"log"
	"wterm/pkg/api"
	"wterm/pkg/weather"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	weatherData, err := api.GetDailyWeatherData("52.191570", "-1.706990")

	if err != nil {
		log.Fatalln(err)
	} else {
		weather.PrintDailyTable(&weatherData)
	}
}


