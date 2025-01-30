package quotacounting_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uudashr/go-playground/quotacounting"
)

// TODO: Test for multiple users

func TestQuotaGuard(t *testing.T) {
	cal := newCalendar()
	gw := quotacounting.NewGateway(
		2,   // Daily quota
		5,   // Weekly quota
		cal, // Calendar
	)

	// Day 1
	cal.setDay(1)
	err := gw.Send("user-01", "msg-01")
	require.NoError(t, err)

	err = gw.Send("user-01", "msg-02")
	require.NoError(t, err)

	err = gw.Send("user-01", "msg-03")
	require.ErrorIs(t, err, quotacounting.ErrDailyQuotaExceeded)

	// Day 2
	cal.setDay(2)
	err = gw.Send("user-01", "msg-04")
	require.NoError(t, err)

	// Day 3
	cal.setDay(3)
	err = gw.Send("user-01", "msg-05")
	require.NoError(t, err)

	err = gw.Send("user-01", "msg-06")
	require.NoError(t, err)

	// Day 4
	cal.setDay(4)
	err = gw.Send("user-01", "msg-07")
	require.ErrorIs(t, err, quotacounting.ErrWeeklyQuotaExceeded)
}

type calendar struct {
	day  int
	week int
}

func newCalendar() *calendar {
	return &calendar{
		day:  1,
		week: 1,
	}
}

func (c *calendar) setDay(d int) {
	if d < 1 {
		panic("day starts from 1")
	}

	c.day = d
	c.week = (d-1)/7 + 1
}

func (c *calendar) Now() (day, week int) {
	return c.day, c.week
}
