package timeutil_test

import (
	"testing"
	"time"

	"github.com/uudashr/go-playground/fmtime/internal/timeutil"
)

func TestDayCompare(t *testing.T) {
	timeLayout := "2006-01-02T15:04:05-07:00"
	tests := []struct {
		name   string
		time1  string
		time2  string
		expect int
	}{
		{
			name:   "same time",
			time1:  "2025-07-25T12:45:30+07:00",
			time2:  "2025-07-25T12:45:30+07:00",
			expect: 0,
		},
		{
			name:   "earlier, same day",
			time1:  "2025-07-25T10:45:30+07:00",
			time2:  "2025-07-25T12:45:30+07:00",
			expect: 0,
		},
		{
			name:   "later, same day",
			time1:  "2025-07-25T12:45:30+07:00",
			time2:  "2025-07-25T15:45:30+07:00",
			expect: 0,
		},
		{
			name:   "day before",
			time1:  "2025-07-24T10:45:30+07:00",
			time2:  "2025-07-25T12:45:30+07:00",
			expect: -1,
		},
		{
			name:   "day after",
			time1:  "2025-07-26T12:45:30+07:00",
			time2:  "2025-07-25T12:45:30+07:00",
			expect: 1,
		},
		{
			name:   "month before",
			time1:  "2025-06-25T10:45:30+07:00",
			time2:  "2025-07-25T12:45:30+07:00",
			expect: -1,
		},
		{
			name:   "month after",
			time1:  "2025-08-25T12:45:30+07:00",
			time2:  "2025-07-25T12:45:30+07:00",
			expect: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			time1, err := time.ParseInLocation(timeLayout, tt.time1, time.Local)
			if err != nil {
				t.Fatalf("Fail to parse time1 %q: %v", tt.time1, err)
			}

			time2, err := time.ParseInLocation(timeLayout, tt.time2, time.Local)
			if err != nil {
				t.Fatalf("Fail to parse time2 %q: %v", tt.time1, err)
			}

			// Act
			res := timeutil.DayCompare(time1, time2)

			// Assert
			if got, want := res, tt.expect; got != want {
				t.Errorf("DayCompare got: %d, want: %d", got, want)
			}
		})
	}
}

func TestDayDiff(t *testing.T) {
	timeLayout := "2006-01-02T15:04:05-07:00"
	tests := []struct {
		name       string
		time1      string
		time2      string
		expectDiff int
	}{

		{
			name:       "same time",
			time1:      "2025-07-25T12:45:30+07:00",
			time2:      "2025-07-25T12:45:30+07:00",
			expectDiff: 0,
		},
		{
			name:       "earlier, same day",
			time1:      "2025-07-25T10:45:30+07:00",
			time2:      "2025-07-25T12:45:30+07:00",
			expectDiff: 0,
		},
		{
			name:       "later, same day",
			time1:      "2025-07-25T12:45:30+07:00",
			time2:      "2025-07-25T15:45:30+07:00",
			expectDiff: 0,
		},
		{
			name:       "a day before",
			time1:      "2025-07-24T10:45:30+07:00",
			time2:      "2025-07-25T12:45:30+07:00",
			expectDiff: -1,
		},
		{
			name:       "a day after",
			time1:      "2025-07-26T12:45:30+07:00",
			time2:      "2025-07-25T12:45:30+07:00",
			expectDiff: 1,
		},
		{
			name:       "2 days before",
			time1:      "2025-07-23T10:45:30+07:00",
			time2:      "2025-07-25T12:45:30+07:00",
			expectDiff: -2,
		},
		{
			name:       "2 days after",
			time1:      "2025-07-27T12:45:30+07:00",
			time2:      "2025-07-25T12:45:30+07:00",
			expectDiff: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			time1, err := time.ParseInLocation(timeLayout, tt.time1, time.Local)
			if err != nil {
				t.Fatalf("Fail to parse time1 %q: %v", tt.time1, err)
			}

			time2, err := time.ParseInLocation(timeLayout, tt.time2, time.Local)
			if err != nil {
				t.Fatalf("Fail to parse time2 %q: %v", tt.time2, err)
			}

			// Act
			d := timeutil.DayDiff(time1, time2)

			// Assert
			if got, want := d, tt.expectDiff; got != want {
				t.Errorf("got: %d, want: %d", got, want)
			}
		})
	}
}
