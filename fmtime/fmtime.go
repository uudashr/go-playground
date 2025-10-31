package fmtime

import (
	"time"

	"github.com/uudashr/go-playground/fmtime/internal/timeutil"
)

type Clock interface {
	Now() time.Time
}

type Format struct {
	Clock Clock
}

func (f *Format) Format(t time.Time) string {
	clock := f.Clock
	if clock == nil {
		clock = timeutil.SystemClock
	}

	now := clock.Now()
	c := timeutil.DayDiff(t, now)
	switch c {
	case 0:
		return "today"
	case -1:
		return "yesterday"
	case 1:
		return "tomorrow"
	}

	return t.Format("2 January 2006")
}
