package date

import "time"

func DateForDay(year int, day int) time.Time {
	return time.Date(year, 12, day, 0, 0, 0, 0, time.Local)
}

func DateStringForDay(year int, day int) string {
	return DateForDay(year, day).Format("02.01.2006")
}

func CurrentDay() int {
	return time.Now().Day()
}

func CurrentYear() int {
	return time.Now().Year()
}
