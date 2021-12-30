package weather

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

type WeatherResponse struct {
	Lat             float32
	Lon             float32
	Timezone        string
	Timezone_offset int
	Daily           []Daily
}

const layoutUK = "25 August 1990"

func PrintDailyTable(weatherData *WeatherResponse) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Day"},
			{Align: simpletable.AlignCenter, Text: "Icon"},
			{Align: simpletable.AlignCenter, Text: "Main"},
			{Align: simpletable.AlignCenter, Text: "Description"},
		},
	}
	
	fmt.Printf("Date: %s\n", GetDate(&weatherData.Daily[0]))
	for _, day := range weatherData.Daily {
		weekDay := GetWeekDay(&day)

		for _, daily := range day.Weather {
			emoji := GetIcon(&daily)
			row := []*simpletable.Cell{
				{Text: weekDay},
				{Text: emoji, Align: simpletable.AlignCenter},
				{Text: daily.Main},
				{Text: daily.Description},
			}

			table.Body.Cells = append(table.Body.Cells, row)
		}
	}

	fmt.Println(table.String())
}