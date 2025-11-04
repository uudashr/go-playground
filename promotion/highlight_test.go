package promotion_test

import (
	"testing"
	"time"

	"github.com/uudashr/go-playground/promotion"
)

func TestHighlightBuilder(t *testing.T) {
	timeLayout := "2006-01-02T15:04:05-07:00"
	tests := []struct {
		name                string
		now                 string
		promoExpiresAt      string
		noExpiration        bool
		promoMaxUsage       int
		promoUsageCount     int
		expiresTodayMessage string
		dateFormat          promotion.TimeFormatter
		expirationFormat    string
		utilizationFormat   string
		expectLine1         string
		expectLine2         string
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
			name:                "Case 1, custom expiration format",
			now:                 "2025-01-15T14:45:05+07:00",
			promoExpiresAt:      "2025-01-15T22:00:00+07:00",
			promoMaxUsage:       0,
			expiresTodayMessage: "For today only",
			expectLine1:         "For today only",
			expectLine2:         "Unlimited rides",
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
			name:             "Case 2, custom expires format",
			now:              "2025-01-15T14:45:05+07:00",
			promoExpiresAt:   "2025-03-14T22:00:00+07:00",
			promoUsageCount:  5,
			promoMaxUsage:    15,
			dateFormat:       promotion.TimeFormat("January 2, 2006"),
			expirationFormat: "Until %s",
			expectLine1:      "Until March 14, 2025",
			expectLine2:      "10 rides left",
		},
		{
			name:              "Case 2, custom utilization",
			now:               "2025-01-15T14:45:05+07:00",
			promoExpiresAt:    "2025-03-14T22:00:00+07:00",
			promoUsageCount:   5,
			promoMaxUsage:     15,
			utilizationFormat: "%d rides remaining",
			expectLine1:       "Expires 14 March 2025",
			expectLine2:       "10 rides remaining",
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
				t.Fatalf("Failed to parse 'now' %q: %v", tt.promoExpiresAt, err)
			}

			var expiresAt time.Time
			if tt.promoExpiresAt != "" {
				expiresAt, err = time.ParseInLocation(timeLayout, tt.promoExpiresAt, time.Local)
				if err != nil {
					t.Fatalf("Failed to parse 'promoExpiresAt' %q: %v", tt.promoExpiresAt, err)
				}
			}

			builder := &promotion.HighlightBuilder{
				Expiration: promotion.ExpirationFormat{
					ExpiresTodayMessage: tt.expiresTodayMessage,
					DateFormat:          tt.dateFormat,
					Format:              tt.expirationFormat,
				},
				Utilization: promotion.UtilizationFormat{
					Format: tt.utilizationFormat,
				},
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
				t.Fatalf("Failed to build highlight: %v", err)
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
				t.Fatalf("Failed to parse 'now' %q: %v", tt.now, err)
			}

			expiresAt, err := time.ParseInLocation(timeLayout, tt.expiresAt, time.Local)
			if err != nil {
				t.Fatalf("Failed to parse 'expiresAt' %q: %v", tt.expiresAt, err)
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

func TestExpirationFormat(t *testing.T) {
	timeLayout := "2006-01-02T15:04:05-07:00"
	tests := []struct {
		name                   string
		now                    string
		expiresAt              string
		dateFormat             promotion.TimeFormatter
		expirationFormat       string
		expirationTodayMessage string
		expectError            error
		expectExpiration       string
	}{
		{
			name:             "Case 1",
			now:              "2025-01-15T14:45:05+07:00",
			expiresAt:        "2025-01-15T22:00:00+07:00",
			expectExpiration: "Available only today",
		},
		{
			name:                   "Case 1, custom today message",
			now:                    "2025-01-15T14:45:05+07:00",
			expiresAt:              "2025-01-15T22:00:00+07:00",
			expirationTodayMessage: "Until today only",
			expectExpiration:       "Until today only",
		},
		{
			name:             "Case 2",
			now:              "2025-01-15T14:45:05+07:00",
			expiresAt:        "2025-03-14T22:00:00+07:00",
			expectExpiration: "Expires 14 March 2025",
		},
		{
			name:             "Case 2, custom date format",
			now:              "2025-01-15T14:45:05+07:00",
			expiresAt:        "2025-03-14T22:00:00+07:00",
			dateFormat:       promotion.TimeFormat("January 2, 2006"),
			expectExpiration: "Expires March 14, 2025",
		},
		{
			name:             "Case 2, ID language",
			now:              "2025-01-15T14:45:05+07:00",
			expiresAt:        "2025-03-14T22:00:00+07:00",
			expirationFormat: "Sampai dengan %s",
			dateFormat: &promotion.SimpleDateFormat{
				Months: []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "July", "Agustus", "September", "Oktober", "November", "Desember"},
			},
			expectExpiration: "Sampai dengan 14 Maret 2025",
		},
		{
			name:             "Case 2, custom expiration format",
			now:              "2025-01-15T14:45:05+07:00",
			expiresAt:        "2025-03-14T22:00:00+07:00",
			expirationFormat: "Expires at %s",
			expectExpiration: "Expires at 14 March 2025",
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
				t.Fatalf("Failed to parse 'now' %q: %v", tt.now, err)
			}

			expiresAt, err := time.ParseInLocation(timeLayout, tt.expiresAt, time.Local)
			if err != nil {
				t.Fatalf("Failed to parse 'expiresAt' %q: %v", tt.expiresAt, err)
			}

			fmt := promotion.ExpirationFormat{
				ExpiresTodayMessage: tt.expirationTodayMessage,
				DateFormat:          tt.dateFormat,
				Format:              tt.expirationFormat,
			}

			// Act
			out, err := fmt.FormatExpiration(now, expiresAt)

			// Assert
			if got, want := err, tt.expectError; got != want {
				t.Fatalf("Format err got: %v, want: %v", got, want)
			}

			if got, want := out, tt.expectExpiration; got != want {
				t.Errorf("Format got: %q, want: %q", got, want)
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

func TestUtilizationFormat(t *testing.T) {
	tests := []struct {
		name              string
		maxUsage          int
		usageCount        int
		unlimitedMessage  string
		utilizationFormat string
		expectError       error
		expectUtilization string
	}{
		{
			name:              "Case 1",
			maxUsage:          0,
			expectUtilization: "Unlimited rides",
		},
		{
			name:              "Case 1, custom unlimited message",
			maxUsage:          0,
			unlimitedMessage:  "No limit",
			expectUtilization: "No limit",
		},
		{
			name:              "Case 2",
			usageCount:        5,
			maxUsage:          15,
			expectUtilization: "10 rides left",
		},
		{
			name:              "Case 2, custom format",
			usageCount:        5,
			maxUsage:          15,
			utilizationFormat: "%d rides remaining",
			expectUtilization: "10 rides remaining",
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
			// Arrange
			fmt := promotion.UtilizationFormat{
				UnlimitedMessage: tt.unlimitedMessage,
				Format:           tt.utilizationFormat,
			}

			// Act
			out, err := fmt.FormatUtilization(tt.usageCount, tt.maxUsage)

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

func TestSimpleDateFormat(t *testing.T) {
	timeLayout := "2006-01-02T15:04:05-07:00"
	tests := []struct {
		name      string
		time      string
		months    []string
		expectOut string
	}{
		{
			name:      "No month defined",
			time:      "2025-01-15T14:45:05+07:00",
			expectOut: "15 January 2025",
		},
		{
			name:      "With month defined",
			time:      "2025-01-15T14:45:05+07:00",
			months:    []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "July", "Agustus", "September", "Oktober", "November", "Desember"},
			expectOut: "15 Januari 2025",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, err := time.ParseInLocation(timeLayout, tt.time, time.Local)
			if err != nil {
				t.Fatalf("Failed to parse 'time' %q: %v", tt.time, err)
			}

			// Arrange
			df := promotion.SimpleDateFormat{
				Months: tt.months,
			}

			// Act
			out := df.Format(val)

			// Assert
			if got, want := out, tt.expectOut; got != want {
				t.Errorf("Format got: %q, want: %q", got, want)
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
