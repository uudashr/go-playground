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

func FormatExpiration(now time.Time, expiresAt time.Time) (string, error) {
	days := cal.DayDiff(expiresAt, now)
	if days < 0 {
		return "", ErrExpired
	}

	// Expiration can happen on the same day
	if !now.Before(expiresAt) {
		return "", ErrExpired
	}

	if days == 0 {
		return "Available only today", nil
	}

	return "Expires " + expiresAt.Format("2 January 2006"), nil
}

var (
	ErrOverused   = errors.New("overused error")
	ErrReachedMax = errors.New("reached max error")
)

func FormatUtilization(usageCount int, maxUsage int) (string, error) {
	if maxUsage == 0 {
		return "Unlimited rides", nil
	}

	remaining := maxUsage - usageCount
	if remaining == 0 {
		return "", ErrReachedMax
	}

	if remaining < 0 {
		return "", ErrOverused
	}

	return fmt.Sprintf("%d rides left", remaining), nil
}
