package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"wterm/pkg/weather"
)


func makeURL(lat string, long string) string {
	var API_KEY string = os.Getenv("API_KEY")

	exclude := "current,minutely,hourly,alerts"
	const BASE_URL string = "https://api.openweathermap.org/data/2.5/onecall?lat=%v&lon=%v&exclude=%v&appid=%v"
	return fmt.Sprintf(BASE_URL, lat, long, exclude, API_KEY)
}

func GetDailyWeatherData(lat string, long string) (weather.WeatherResponse, error) {
	var weatherData weather.WeatherResponse
	url := makeURL(lat, long)

	resp, err := http.Get(url)
	if err != nil {
		return weatherData, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return weatherData, err
	}

	weatherJson := string(body)
	jsonErr := json.Unmarshal([]byte(weatherJson), &weatherData)

	if jsonErr != nil {
		return weatherData, jsonErr
	}

	return weatherData, nil
}