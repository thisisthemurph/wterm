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

func (d *Daily) GetWeekDay() string {
	return time.Unix(d.Dt, 0).Weekday().String()
}

func (d *Daily) GetDate() string {
	t := time.Unix(d.Dt, 0)
	return t.Format("Monday - Jan _2 2006")
}