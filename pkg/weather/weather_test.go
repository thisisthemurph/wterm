package weather

import "testing"

func TestGetIcon(t *testing.T) {
	t.Run("various icon codes", func(t *testing.T) {
		weatherItems := weatherFactory()
		wantItems := [5]string{"🌞", "🌞", "🌧️", "❄️", "🌫️"}
	
		for index, weather := range weatherItems {
			got := GetIcon(&weather)
			want := wantItems[index]
			assertIcon(t, got, want)
		}
	})

	t.Run("unknown icon code", func(t *testing.T) {
		weather := makeWeather("00d")

		got := GetIcon(&weather)
		want := "❔"
		assertIcon(t, got, want)
	})

	t.Run("empty stringcode", func(t *testing.T) {
		weather := makeWeather("")

		got := GetIcon(&weather)
		want := "❔"
		assertIcon(t, got, want)
	})
}

func makeWeather(icon string) Weather {
	return Weather{
		500,
		"Rain",
		"slight rain",
		icon,
	}
}

func weatherFactory() []Weather {
	codes := []string{"01d", "01n", "09d", "13d", "50n"}
	
	var wetherItems []Weather
	for _, code := range codes {
		wetherItems = append(wetherItems, makeWeather(code))
	}

	return wetherItems
}

func assertIcon(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}