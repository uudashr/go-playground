package promotion_test

import (
	"testing"
	"time"

	"github.com/uudashr/go-playground/promotion"
)

func TestHighlight(t *testing.T) {
	timeLayout := "2006-01-02T15:04:05-07:00"
	tests := []struct {
		name            string
		now             string
		promoExpiresAt  string
		noExpiration    bool
		promoMaxUsage   int
		promoUsageCount int
		expectLine1     string
		expectLine2     string
	}{
		{
			name:           "Case 1",
			now:            "2025-01-15T14:45:05+07:00",
			promoExpiresAt: "2025-01-15T22:00:00+07:00",
			promoMaxUsage:  0,
			expectLine1:    "Available only today",
			expectLine2:    "Unlimited rides",
		},
		{
			name:            "Case 2",
			now:             "2025-01-15T14:45:05+07:00",
			promoExpiresAt:  "2025-03-14T22:00:00+07:00",
			promoUsageCount: 5,
			promoMaxUsage:   15,
			expectLine1:     "Expires 14 March 2025",
			expectLine2:     "10 rides left",
		},
		{
			name:          "No expiration",
			now:           "2025-01-15T14:45:05+07:00",
			noExpiration:  true,
			promoMaxUsage: 0,
			expectLine1:   "",
			expectLine2:   "Unlimited rides",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			now, err := time.ParseInLocation(timeLayout, tt.now, time.Local)
			if err != nil {
				t.Fatalf("Fail to parse now %q: %v", tt.promoExpiresAt, err)
			}

			var expiresAt time.Time
			if tt.promoExpiresAt != "" {
				expiresAt, err = time.ParseInLocation(timeLayout, tt.promoExpiresAt, time.Local)
				if err != nil {
					t.Fatalf("Fail to parse promoExpiresAt %q: %v", tt.promoExpiresAt, err)
				}
			}

			builder := &promotion.HighlightBuilder{
				Clock: &staticClock{now},
			}

			// Act
			highlight, err := builder.Build(promotion.HighlightSpec{
				ExpiresAt:    expiresAt,
				NoExpiration: tt.noExpiration,
				UsageCount:   tt.promoUsageCount,
				MaxUsage:     tt.promoMaxUsage,
			})
			if err != nil {
				t.Fatalf("Fail to build highlight: %v", err)
			}

			// Assert
			if got, want := highlight.Expiration, tt.expectLine1; got != want {
				t.Errorf("Expiration got: %q, want: %q", got, want)
			}

			if got, want := highlight.Utilization, tt.expectLine2; got != want {
				t.Errorf("Utilization got: %q, want: %q", got, want)
			}
		})
	}
}

func TestFormatExpiration(t *testing.T) {
	timeLayout := "2006-01-02T15:04:05-07:00"
	tests := []struct {
		name             string
		now              string
		expiresAt        string
		expectError      error
		expectExpiration string
	}{
		{
			name:             "Case 1",
			now:              "2025-01-15T14:45:05+07:00",
			expiresAt:        "2025-01-15T22:00:00+07:00",
			expectExpiration: "Available only today",
		},
		{
			name:             "Case 2",
			now:              "2025-01-15T14:45:05+07:00",
			expiresAt:        "2025-03-14T22:00:00+07:00",
			expectExpiration: "Expires 14 March 2025",
		},
		{
			name:        "Reached expiration on the same day",
			now:         "2025-01-15T14:45:05+07:00",
			expiresAt:   "2025-01-15T13:45:05+07:00",
			expectError: promotion.ErrExpired,
		},
		{
			name:        "Reached expiration on previous day",
			now:         "2025-01-15T14:45:05+07:00",
			expiresAt:   "2025-01-14T14:45:05+07:00",
			expectError: promotion.ErrExpired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			now, err := time.ParseInLocation(timeLayout, tt.now, time.Local)
			if err != nil {
				t.Fatalf("Fail to parse now %q: %v", tt.now, err)
			}

			expiresAt, err := time.ParseInLocation(timeLayout, tt.expiresAt, time.Local)
			if err != nil {
				t.Fatalf("Fail to parse expiresAt %q: %v", tt.expiresAt, err)
			}

			// Act
			out, err := promotion.FormatExpiration(now, expiresAt)

			// Assert
			if got, want := err, tt.expectError; got != want {
				t.Fatalf("FormatExpiration err got: %v, want: %v", got, want)
			}

			if got, want := out, tt.expectExpiration; got != want {
				t.Errorf("FormatExpiration got: %q, want: %q", got, want)
			}
		})
	}
}

func TestFormatUtilization(t *testing.T) {
	tests := []struct {
		name              string
		maxUsage          int
		usageCount        int
		expectError       error
		expectUtilization string
	}{
		{
			name:              "Case 1",
			maxUsage:          0,
			expectUtilization: "Unlimited rides",
		},
		{
			name:              "Case 2",
			usageCount:        5,
			maxUsage:          15,
			expectUtilization: "10 rides left",
		},
		{
			name:              "No max, has some usage",
			maxUsage:          0,
			usageCount:        2,
			expectUtilization: "Unlimited rides",
		},
		{
			name:        "Overused",
			maxUsage:    3,
			usageCount:  4,
			expectError: promotion.ErrOverused,
		},
		{
			name:        "Reached max",
			maxUsage:    3,
			usageCount:  3,
			expectError: promotion.ErrReachedMax,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			out, err := promotion.FormatUtilization(tt.usageCount, tt.maxUsage)

			// Assert
			if got, want := err, tt.expectError; got != want {
				t.Errorf("FormatUtilization err got: %v, want: %v", got, want)
			}

			if got, want := out, tt.expectUtilization; got != want {
				t.Errorf("FormatUtilization got: %q, want: %q", got, want)
			}

		})
	}
}

type staticClock struct {
	now time.Time
}

func (sc *staticClock) Now() time.Time {
	return sc.now
}
