package weather

import (
	"strconv"
)

type Weather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

func (w *Weather) GetIcon() string {
	code := iconCodeToInt(w.Icon)

	switch code {
		case 1:
			return "đ"
		case 2:
			return "đ¤ī¸"
		case 3, 4:
			return "âī¸"
		case 9:
			return "đ§ī¸"
		case 10:
			return "đĻī¸"
		case 11:
			return "đŠī¸"
		case 13:
			return "âī¸"
		case 50:
			return "đĢī¸"
		default:
			return "â"
	}
}

func iconCodeToInt(iconCode string) int64 {
	if len(iconCode) != 3 {
		return -1
	}	

	code, err := strconv.ParseInt(iconCode[0:2], 10, 32)
	if err != nil {
		return -1
	}

	return code
}