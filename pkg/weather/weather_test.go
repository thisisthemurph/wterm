package weather_test

import (
	"testing"
	weather "wterm/pkg/weather"
)

func TestGetIcon(t *testing.T) {
	t.Run("various icon codes", func(t *testing.T) {
		weatherItems := weatherFactory()
		wantItems := [5]string{"🌞", "🌞", "🌧️", "❄️", "🌫️"}
	
		for index, weather := range weatherItems {
			got := weather.GetIcon()
			want := wantItems[index]
			assertIcon(t, got, want)
		}
	})

	t.Run("unknown icon code", func(t *testing.T) {
		weather := makeWeather("00d")

		got := weather.GetIcon()
		want := "❔"
		assertIcon(t, got, want)
	})

	t.Run("empty stringcode", func(t *testing.T) {
		weather := makeWeather("")

		got := weather.GetIcon()
		want := "❔"
		assertIcon(t, got, want)
	})
}

func makeWeather(iconCode string) weather.Weather {
	return weather.Weather{
		500,
		"Rain",
		"slight rain",
		iconCode,
	}
}

func weatherFactory() []weather.Weather {
	codes := []string{"01d", "01n", "09d", "13d", "50n"}
	
	var wetherItems []weather.Weather
	for _, iconCode := range codes {
		wetherItems = append(wetherItems, makeWeather(iconCode))
	}

	return wetherItems
}

func assertIcon(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}