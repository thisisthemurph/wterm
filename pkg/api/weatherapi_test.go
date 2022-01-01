package api

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestGetDailyWeatherData(t *testing.T) {
	godotenv.Load("../../.env")

	t.Run("normal request", func(t *testing.T) {
		got, err := GetDailyWeatherData("27.994402", "-81.760254")
		gotLen := len(got.Daily)
		wantLen := 8
		
		assertNoError(t, err)
		assertSliceLength(t, gotLen, wantLen)
	})
	
	t.Run("bad coordinates request", func(t *testing.T) {
		_, err := GetDailyWeatherData("BAD_LAT", "BAD_LONG")
		assertHasError(t, err)
	})

	t.Run("out of bounds coordinates", func(t *testing.T) {
		_, err := GetDailyWeatherData("-91", "181")
		assertHasError(t, err)
	})

	t.Run("empty coordinate strings", func(t *testing.T) {
		_, err := GetDailyWeatherData("", "")
		assertHasError(t, err)
	})
}

func assertHasError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("There was no error when an error was expected.")
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("There was an error when one was not expected. Error: '%s'", err.Error())
	}
}

func assertSliceLength(t *testing.T, gotLen int, wantLen int) {
	t.Helper()
	if gotLen != wantLen {
		t.Errorf("got slice length of %d wantd length %d", gotLen, wantLen)
	}
}