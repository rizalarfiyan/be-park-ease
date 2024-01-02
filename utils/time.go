package utils

import "time"

func RemaniningToday() time.Duration {
	now := time.Now()
	dateTomorrow := now.AddDate(0, 0, 1)
	tomorrow := time.Date(dateTomorrow.Year(), dateTomorrow.Month(), dateTomorrow.Day(), 0, 0, 0, 0, dateTomorrow.Location())
	return tomorrow.Sub(now)
}

func StartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func EndOfDay(date time.Time) time.Time {
	return StartOfDay(date).AddDate(0, 0, 1).Add(-time.Millisecond)
}

func StartOfWeek(date time.Time) time.Time {
	firstDayOfWeek := date.AddDate(0, 0, -int(date.Weekday()-time.Sunday))
	return StartOfDay(firstDayOfWeek)
}

func EndOfWeek(date time.Time) time.Time {
	endDayOfWeek := date.AddDate(0, 0, int(time.Saturday-date.Weekday()))
	return EndOfDay(endDayOfWeek)
}

func StartOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

func EndOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month()+1, 1, 0, 0, 0, 0, date.Location()).Add(-time.Millisecond)
}
