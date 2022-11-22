package utils

import "time"

func DateForDay(year int, day int) time.Time {
	return time.Date(year, 12, day, 0, 0, 0, 0, time.Local)
}

func DateStringForDay(year int, day int) string {
	return DateForDay(year, day).Format("02.01.2006")
}
