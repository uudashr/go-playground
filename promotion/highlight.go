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
	ExpiresAt  time.Time
	UsageCount int
	MaxUsage   int
}

type HighlightBuilder struct {
	Clock Clock
}

func (hb *HighlightBuilder) Build(spec HighlightSpec) (*Highlight, error) {
	now := hb.clock().Now()
	exp, err := FormatExpiration(now, spec.ExpiresAt)
	if err != nil {
		return nil, err
	}

	utilization, err := FormatUtilization(spec.UsageCount, spec.MaxUsage)
	if err != nil {
		return nil, err
	}

	return &Highlight{
		Expiration:  exp,
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

func FormatExpiration(now time.Time, expiresAt time.Time) (string, error) {
	days := cal.DayDiff(expiresAt, now)
	if days < 0 {
		return "", errors.New("expired")
	}

	if days == 0 {
		return "Available only today", nil
	}

	return "Expires " + expiresAt.Format("2 January 2006"), nil
}

func FormatUtilization(usageCount int, maxUsage int) (string, error) {
	if maxUsage == 0 {
		return "Unlimited rides", nil
	}

	remaining := maxUsage - usageCount
	if remaining <= 0 {
		return "", fmt.Errorf("invalid remaining %d", remaining)
	}

	return fmt.Sprintf("%d rides left", remaining), nil
}
