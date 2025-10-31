package timeutil

import (
	"time"
)

func DayCompare(a time.Time, b time.Time) int {
	sameDay := a.Year() == b.Year() && a.Month() == b.Month() && a.Day() == b.Day()
	if sameDay {
		return 0
	}

	if a.Before(b) {
		return -1
	}

	return 1
}

func DayDiff(a time.Time, b time.Time) int {
	a = DayTruncate(a)
	b = DayTruncate(b)
	return int(a.Sub(b).Hours() / 24)
}

func DayTruncate(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

var SystemClock = &systemClock{}

type systemClock struct {
}

func (sc systemClock) Now() time.Time {
	return time.Now()
}
