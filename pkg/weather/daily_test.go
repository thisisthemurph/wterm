package weather

import "testing"

func makeDaily() Daily {
	temp := Temperature{132.3, 132.3, 132.3, 132.3, 132.3, 132.3}
	w := Weather{500, "Rain", "slight rain", "09d"}
	weather := []Weather{w}

	return Daily{
		1640955380,
		1640955380,
		1640955380,
		1640955380,
		0.86,
		temp,
		weather,
	}
}

var testDaily Daily = makeDaily()

func TestGetWeekDay(t *testing.T) {
	got := GetWeekDay(&testDaily)
	want := "Friday"

	if got != want {
		t.Errorf("got %v want %s", got, want)
	}
}

func TestGetDate(t *testing.T) {
	got := GetDate(&testDaily)
	want := "Friday - Dec 31 2021"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}