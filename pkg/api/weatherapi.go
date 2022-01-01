package api

import (
	"encoding/json"
	"errors"
	"strconv"
	"wterm/pkg/weather"
)

func GetDailyWeatherData(lat string, long string) (weather.WeatherResponse, error) {
	var weatherData weather.WeatherResponse

	// Make the API call
	jsonData, err := client(lat, long)
	if err != nil {
		return weatherData, err
	}

	jsonErr := json.Unmarshal([]byte(jsonData), &weatherData)
	if jsonErr != nil {
		return weatherData, jsonErr
	}

	return weatherData, nil
}

func validateCoordinates(lat string, long string) error {
	latErr := validateLatitude(lat)
	longErr := validateLongitude(long)

	if latErr != nil {
		return latErr
	} else if longErr != nil {
		return longErr
	}

	return nil
}

func validateLatitude(lat string) error {
	latitude, err := strconv.ParseFloat(lat, 64,)

	if err != nil {
		return errors.New("Latitude is not in expected format.")
	}

	if latitude < -90 || latitude > 90 {
		return errors.New("Latitude out of bounds")
	}

	return nil
}

func validateLongitude(long string) error {
	longitude, err := strconv.ParseFloat(long, 64)

	if err != nil {
		return errors.New("Longitude is not in expected format.")
	}

	if longitude < -180 || longitude > 180 {
		return errors.New("Longitude out of bounds.")
	}

	return nil
}