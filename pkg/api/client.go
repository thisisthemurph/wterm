package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func client(lat string, long string) (string, error) {
	err := validateCoordinates(lat, long)
	if err != nil {
		return "", err
	}

	endpoint := makeURL(lat, long)
	resp, err := http.Get(endpoint)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("There has been an issue querying the endpoint.")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func makeURL(lat string, long string) string {
	var API_KEY string = os.Getenv("API_KEY")

	exclude := "current,minutely,hourly,alerts"
	const BASE_URL string = "https://api.openweathermap.org/data/2.5/onecall?lat=%v&lon=%v&exclude=%v&appid=%v"
	return fmt.Sprintf(BASE_URL, lat, long, exclude, API_KEY)
}