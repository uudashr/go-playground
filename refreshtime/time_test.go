package refreshtime_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/uudashr/go-playground/refreshtime"
)

func TestRefreshTime(t *testing.T) {
	currentTime := time.Now()

	tests := []struct {
		expiresIn time.Duration
		expected  time.Duration
	}{
		{10 * time.Minute, time.Duration(float64(10*time.Minute) * 0.2)}, // 20% of 10 minutes
		{30 * time.Minute, 5 * time.Minute},                              // 5 minutes
		{2 * time.Hour, 10 * time.Minute},                                // 10 minutes
	}

	for _, tt := range tests {
		t.Run(tt.expiresIn.String(), func(t *testing.T) {
			expiresAt := currentTime.Add(tt.expiresIn)
			got := refreshtime.RefreshTime(currentTime, tt.expiresIn)
			expected := currentTime.Add(tt.expected)
			require.Equal(t, expected, got)
			require.Less(t, got, expiresAt)
		})
	}
}
