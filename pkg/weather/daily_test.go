package weather_test

import (
	"testing"
	"wterm/pkg/weather"
)

func makeDaily() weather.Daily {
	temp := weather.Temperature{132.3, 132.3, 132.3, 132.3, 132.3, 132.3}
	w := weather.Weather{500, "Rain", "slight rain", "09d"}
	weatherData := []weather.Weather{w}

	return weather.Daily{
		1640955380,
		1640955380,
		1640955380,
		1640955380,
		0.86,
		temp,
		weatherData,
	}
}

var testDaily weather.Daily = makeDaily()

func TestGetWeekDay(t *testing.T) {
	got := testDaily.GetWeekDay()
	want := "Friday"

	if got != want {
		t.Errorf("got %v want %s", got, want)
	}
}

func TestGetDate(t *testing.T) {
	got := testDaily.GetDate()
	want := "Friday - Dec 31 2021"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}