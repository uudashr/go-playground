package fmtime_test

import (
	"testing"
	"time"

	"github.com/uudashr/go-playground/fmtime"
)

func TestFormat(t *testing.T) {
	timeLayout := "2006-01-02T15:04:05-07:00"

	tests := []struct {
		name      string
		now       string
		time      string
		expectOut string
	}{
		{
			name:      "End of today",
			now:       "2025-07-31T12:10:05+07:00",
			time:      "2025-07-31T23:59:59+07:00",
			expectOut: "today",
		},
		{
			name:      "Noon",
			now:       "2025-07-31T12:10:05+07:00",
			time:      "2025-07-31T14:10:10+07:00",
			expectOut: "today",
		},
		{
			name:      "Tomorrow at noon",
			now:       "2025-07-31T12:10:05+07:00",
			time:      "2025-08-01T14:10:10+07:00",
			expectOut: "tomorrow",
		},
		{
			name:      "Yesterday at noon",
			now:       "2025-07-31T12:10:05+07:00",
			time:      "2025-07-30T14:10:10+07:00",
			expectOut: "yesterday",
		},
		{
			name:      "Day after tomorrow",
			now:       "2025-07-31T12:10:05+07:00",
			time:      "2025-08-02T14:10:10+07:00",
			expectOut: "2 August 2025",
		},
		{
			name:      "Next 3 days",
			now:       "2025-07-31T12:10:05+07:00",
			time:      "2025-08-03T14:10:10+07:00",
			expectOut: "3 August 2025",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			now, err := time.ParseInLocation(timeLayout, tt.now, time.Local)
			if err != nil {
				t.Fatalf("Fail to parse now %q: %v", tt.now, err)
			}

			tm, err := time.ParseInLocation(timeLayout, tt.time, time.Local)
			if err != nil {
				t.Fatalf("Fail to parser time %q: %v", tt.time, err)
			}

			tf := &fmtime.Format{
				Clock: &StaticClock{now: now},
			}

			// Act
			out := tf.Format(tm)

			// Assert
			if got, want := out, tt.expectOut; got != want {
				t.Errorf("Format got: %q, want: %q", got, want)
			}
		})
	}
}

type StaticClock struct {
	now time.Time
}

func (sc *StaticClock) Now() time.Time {
	return sc.now
}
