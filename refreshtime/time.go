package refreshtime

import "time"

func RefreshTime(currentTime time.Time, expiresIn time.Duration) time.Time {

	if expiresIn <= 15*time.Minute {
		// short-lived token, buffer would be 20% of expiresIn
		bufferedExp := time.Duration(float64(expiresIn) * 0.2)
		return currentTime.Add(bufferedExp)
	}

	if expiresIn <= 1*time.Hour {
		// long-lived token, buffer would be 5 minutes
		bufferedExp := 5 * time.Minute
		return currentTime.Add(bufferedExp)
	}

	// buffer would be 10 minutes
	return currentTime.Add(10 * time.Minute)
}
