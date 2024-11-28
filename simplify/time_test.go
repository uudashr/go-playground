package simplify_test

import (
	"testing"
	"time"
)

func TestUntil(t *testing.T) {
	expiry := time.Now().Add(2 * time.Hour)
	d1 := expiry.Sub(time.Now())
	d2 := time.Until(expiry)

	if got, want := d1, d2; !equalDuration(got, want, time.Millisecond) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestUntilNegative(t *testing.T) {
	expiry := time.Now().Add(-2 * time.Hour)
	d1 := expiry.Sub(time.Now())
	d2 := time.Until(expiry)

	if got, want := d1, d2; !equalDuration(got, want, time.Millisecond) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestUntilZero(t *testing.T) {
	expiry := time.Now()
	d1 := expiry.Sub(time.Now())
	d2 := time.Until(expiry)

	if got, want := d1, d2; !equalDuration(got, want, time.Millisecond) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestUntilMinutes(t *testing.T) {
	expiry := time.Now().Add(30 * time.Minute)
	d1 := expiry.Sub(time.Now())
	d2 := time.Until(expiry)

	if got, want := d1, d2; !equalDuration(got, want, time.Millisecond) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestUntilSeconds(t *testing.T) {
	expiry := time.Now().Add(45 * time.Second)
	d1 := expiry.Sub(time.Now())
	d2 := time.Until(expiry)

	if got, want := d1, d2; !equalDuration(got, want, time.Millisecond) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func equalDuration(got, want time.Duration, tolerance time.Duration) bool {
	return got >= want-tolerance && got <= want+tolerance
}
