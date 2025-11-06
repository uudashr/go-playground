package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

func main() {
	// Fixed 100ms
	fixed := DelayFixed(100 * time.Millisecond)
	fmt.Printf("fixed: %v\n", fixed)

	// Uniform 50-150ms
	// uniform, err := DelayUniformFromMS(50, 150)
	uniform, err := DelayUniform(50*time.Millisecond, 150*time.Millisecond)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Uniform: %v\n", uniform)

	// Lognormal with median=100ms, sigma=0.4
	// lognormal, err := DelayLognormalMS(100, 0.4)
	lognormal, err := DelayLognormal(100*time.Millisecond, 0.4)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Lognormal: %v\n", lognormal)

	// Lognormal with cap at 1s to tame outliers
	// cappedLognormal, err := DelayLognormalCappedMS(100, 0.4, 1000)
	cappedLognormal, err := DelayLognormalCapped(100*time.Millisecond, 0.4, 1000*time.Millisecond)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Lognormal with cap: %v\n", cappedLognormal)

}

func DelayFixed(d time.Duration) time.Duration {
	return d
}

func DelayUniformFromMS(minMS, maxMS float64) (time.Duration, error) {
	if maxMS <= minMS {
		return 0, errors.New("maxMS must be > minMS")
	}

	u := rand.Float64() // U ~ Uniform(0,1)
	ms := minMS + (maxMS-minMS)*u

	return time.Duration(ms * float64(time.Millisecond)), nil
}

func DelayUniform(min, max time.Duration) (time.Duration, error) {
	if max <= min {
		return 0, errors.New("max must be > min")
	}

	u := rand.Float64() // U ~ Uniform(0,1)
	d := float64(min) + float64(max-min)*u

	return time.Duration(d), nil
}

func DelayLognormalMS(medianMS, sigma float64) (time.Duration, error) {
	if medianMS <= 0 || sigma <= 0 {
		return 0, errors.New("medianMS and sigma must be > 0")
	}

	mu := math.Log(medianMS)     // median m => mu = ln(m)
	z := rand.NormFloat64()      // Z ~ N(0,1)
	ms := math.Exp(mu + sigma*z) // X = exp(mu + sigma*Z)

	return time.Duration(ms * float64(time.Millisecond)), nil
}

func DelayLognormal(median time.Duration, sigma float64) (time.Duration, error) {
	if median <= 0 || sigma <= 0 {
		return 0, errors.New("median and sigma must be > 0")
	}

	mu := math.Log(float64(median)) // median m => mu = ln(m)
	z := rand.NormFloat64()         // Z ~ N(0,1)
	d := math.Exp(mu + sigma*z)     // X = exp(mu + sigma*Z)

	return time.Duration(d), nil
}

func DelayLognormalCappedMS(medianMS, sigma, capMS float64) (time.Duration, error) {
	if capMS <= 0 {
		return 0, errors.New("capMS must be > 0")
	}

	mu := math.Log(medianMS)
	z := rand.NormFloat64()
	ms := math.Min(math.Exp(mu+sigma*z), capMS)

	return time.Duration(ms * float64(time.Millisecond)), nil
}

func DelayLognormalCapped(median time.Duration, sigma float64, cap time.Duration) (time.Duration, error) {
	if cap <= 0 {
		return 0, errors.New("cap must be > 0")
	}

	mu := math.Log(float64(median))
	z := rand.NormFloat64()
	d := math.Min(math.Exp(mu+sigma*z), float64(cap))

	return time.Duration(d), nil
}
