package quotacounting

import "errors"

var (
	ErrDailyQuotaExceeded  = errors.New("daily quota exceeded")
	ErrWeeklyQuotaExceeded = errors.New("weekly quota exceeded")
)

type Gateway struct {
	dailyQuota  int
	weeklyQuota int
	cal         Calendar

	dayCounts  map[int]int
	weekCounts map[int]int
}

func NewGateway(dailyQuota, weeklyQuota int, cal Calendar) *Gateway {
	return &Gateway{
		dailyQuota:  dailyQuota,
		weeklyQuota: weeklyQuota,
		cal:         cal,
		dayCounts:   make(map[int]int),
		weekCounts:  make(map[int]int),
	}
}

func (g *Gateway) Send(user, msg string) error {
	day, week := g.cal.Now()

	dayCount := g.dayCounts[day]
	weekCount := g.weekCounts[week]

	if dayCount >= g.dailyQuota {
		return ErrDailyQuotaExceeded
	}

	if weekCount >= g.weeklyQuota {
		return ErrWeeklyQuotaExceeded
	}

	g.dayCounts[day] = dayCount + 1
	g.weekCounts[week] = weekCount + 1
	return nil
}

type Calendar interface {
	Now() (day, week int)
}
