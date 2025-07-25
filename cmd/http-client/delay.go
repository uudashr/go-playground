package main

import "fmt"

type delayBehavior int

const (
	delayBehaviorConstant delayBehavior = iota
	delayBehaviorIncreasing
	delayBehaviorDecreasing
)

var delayNames = [...]string{
	"constant",
	"increasing",
	"decreasing",
}

func parseDelayBehavior(s string) (delayBehavior, error) {
	switch s {
	case "constant":
		return delayBehaviorConstant, nil
	case "increasing":
		return delayBehaviorIncreasing, nil
	case "decreasing":
		return delayBehaviorDecreasing, nil
	default:
		return -1, fmt.Errorf("unknown delay type: %s", s)
	}
}

func (d delayBehavior) MarshalText() ([]byte, error) {
	if d < delayBehaviorConstant || d > delayBehaviorDecreasing {
		return nil, fmt.Errorf("unknown delay type: %d", d)
	}

	return []byte(d.String()), nil
}

func (d *delayBehavior) UnmarshalText(text []byte) error {
	val, err := parseDelayBehavior(string(text))
	if err != nil {
		return fmt.Errorf("unmarshal delay type: %w", err)
	}

	*d = val

	return nil
}

func (d delayBehavior) String() string {
	if d < delayBehaviorConstant || d > delayBehaviorDecreasing {
		return "unknown"
	}

	return delayNames[d]
}
