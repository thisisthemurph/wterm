package weather

import "time"

type Daily struct {
	Dt         int64
	Sunrise    int64
	Sunset     int64
	Moonrise   int64
	Moon_phase float32
	Temp       Temperature
	Weather    []Weather
}

func GetWeekDay(dailWeather *Daily) string {
	return time.Unix(dailWeather.Dt, 0).Weekday().String()
}

func GetDate(dailyWeather *Daily) string {
	t := time.Unix(dailyWeather.Dt, 0)
	return t.Format("Monday - Jan _2 2006")
}