package cal

import (
	"time"
)

type Cal struct {
	Clock Clock
}

func (c *Cal) Dist(t time.Time) int {
	now := c.clock().Now()
	return DayDiff(t, now)
}

func (c *Cal) clock() Clock {
	if c.Clock == nil {
		return defaultClock
	}

	return c.Clock
}

func DayDiff(a time.Time, b time.Time) int {
	a = dayTruncate(a)
	b = dayTruncate(b)
	return int(a.Sub(b).Hours() / 24)
}

func dayTruncate(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

type Clock interface {
	Now() time.Time
}

var defaultClock = &sysClock{}

type sysClock struct {
}

func (sc *sysClock) Now() time.Time {
	return time.Now()
}
