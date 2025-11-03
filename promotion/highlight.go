package promotion

import (
	"errors"
	"fmt"
	"time"

	"github.com/uudashr/go-playground/cal"
)

var defaultClock = &sysClock{}

type Highlight struct {
	Expiration  string
	Utilization string
}

type HighlightSpec struct {
	ExpiresAt    time.Time
	NoExpiration bool
	UsageCount   int
	MaxUsage     int
}

type HighlightBuilder struct {
	Clock Clock
}

func (hb *HighlightBuilder) Build(spec HighlightSpec) (*Highlight, error) {
	now := hb.clock().Now()

	var expiration string
	if !spec.NoExpiration {
		exp, err := FormatExpiration(now, spec.ExpiresAt)
		if err != nil {
			return nil, err
		}

		expiration = exp
	}

	utilization, err := FormatUtilization(spec.UsageCount, spec.MaxUsage)
	if err != nil {
		return nil, err
	}

	return &Highlight{
		Expiration:  expiration,
		Utilization: utilization,
	}, nil
}

func (hb *HighlightBuilder) clock() Clock {
	if hb.Clock == nil {
		return defaultClock
	}

	return hb.Clock
}

type Clock interface {
	Now() time.Time
}

type sysClock struct {
}

func (sc sysClock) Now() time.Time {
	return time.Now()
}

var ErrExpired = errors.New("expired error")

type TimeFormat string

func (tf TimeFormat) Format(t time.Time) string {
	return t.Format(string(tf))
}

type SimpleDateFormat struct {
	MonthNames []string
}

func (sdf *SimpleDateFormat) Format(t time.Time) string {
	monthName := sdf.MonthNames[t.Month()-1]
	return fmt.Sprintf("%d %s %d", t.Day(), monthName, t.Year())
}

var (
	DefaultTimeFormat          = TimeFormat("2 January 2006")
	DefaultExpirationFormat    = "Expires %s"
	DefaultExpiresTodayMessage = "Available only today"
)

type TimeFormatter interface {
	Format(time.Time) string
}

type ExpirationFormat struct {
	DateFormat          TimeFormatter
	Format              string
	ExpiresTodayMessage string
}

func (ef *ExpirationFormat) FormatExpiration(now time.Time, expiresAt time.Time) (string, error) {
	days := cal.DayDiff(expiresAt, now)
	if days < 0 {
		return "", ErrExpired
	}

	// Expiration can happen on the same day
	if !now.Before(expiresAt) {
		return "", ErrExpired
	}

	if days == 0 {
		if ef.ExpiresTodayMessage == "" {
			return DefaultExpiresTodayMessage, nil
		}

		return ef.ExpiresTodayMessage, nil
	}

	// Prepare date format
	dateFormat := ef.DateFormat
	if dateFormat == nil {
		dateFormat = DefaultTimeFormat
	}

	// Prepare expiration format
	expirationFormat := ef.Format
	if expirationFormat == "" {
		expirationFormat = DefaultExpirationFormat
	}

	expDate := dateFormat.Format(expiresAt)
	return fmt.Sprintf(expirationFormat, expDate), nil
}

func FormatExpiration(now time.Time, expiresAt time.Time) (string, error) {
	expFormat := ExpirationFormat{}
	return expFormat.FormatExpiration(now, expiresAt)
}

var (
	ErrOverused   = errors.New("overused error")
	ErrReachedMax = errors.New("reached max error")
)

var (
	DefaultUnlimitedMessage  = "Unlimited rides"
	DefaultUtilizationFormat = "%d rides left"
)

type UtilizationFormat struct {
	UnlimitedMessage string
	Format           string
}

func (uf UtilizationFormat) FormatUtilization(usageCount int, maxUsage int) (string, error) {
	if maxUsage == 0 {
		if uf.UnlimitedMessage == "" {
			return DefaultUnlimitedMessage, nil
		}

		return uf.UnlimitedMessage, nil
	}

	remaining := maxUsage - usageCount
	if remaining == 0 {
		return "", ErrReachedMax
	}

	if remaining < 0 {
		return "", ErrOverused
	}

	utilizationFormat := uf.Format
	if utilizationFormat == "" {
		utilizationFormat = DefaultUtilizationFormat
	}

	return fmt.Sprintf(utilizationFormat, remaining), nil
}

func FormatUtilization(usageCount int, maxUsage int) (string, error) {
	fmt := UtilizationFormat{}
	return fmt.FormatUtilization(usageCount, maxUsage)
}
