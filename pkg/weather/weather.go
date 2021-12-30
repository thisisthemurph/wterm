package weather

import "strconv"

type Weather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

func iconCodeToInt(iconCode string) int64 {
	code, err := strconv.ParseInt(iconCode[0:2], 10, 32)
	if err != nil {
		return -1
	}

	return code
}

func GetIcon(weather *Weather) string {
	code := iconCodeToInt(weather.Icon)

	switch code {
		case 1:
			return "🌞"
		case 2:
			return "🌤️"
		case 3, 4:
			return "☁️"
		case 9:
			return "🌧️"
		case 10:
			return "🌦️"
		case 11:
			return "🌩️"
		case 13:
			return "❄️"
		case 50:
			return "🌫️"
		default:
			return "❔"
	}
}