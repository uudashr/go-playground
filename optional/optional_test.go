package optional_test

import (
	"testing"

	"github.com/uudashr/go-playground/optional"
)

func TestOptional(t *testing.T) {
	testCases := map[string]struct {
		val        any
		optVal     optional.Value[string]
		defaultIn  string
		defaultOut string
		empty      bool
	}{
		"simple value": {
			val:        "Hello",
			optVal:     optional.ValueOf("Hello"),
			defaultIn:  "World",
			defaultOut: "Hello",
		},
		"empty value": {
			val:        "",
			empty:      true,
			defaultIn:  "World",
			defaultOut: "World",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			optVal := tc.optVal
			if got, want := optVal.Empty(), tc.empty; got != want {
				t.Errorf("got %v, want %v", got, want)
			}

			defaultVal := optVal.Default("World")
			if got, want := defaultVal, tc.defaultOut; got != want {
				t.Errorf("got %v, want %v", got, want)
			}

			val, ok := optVal.Get()
			if got, want := ok, !tc.empty; got != want {
				t.Fatalf("got %v, want %v", got, want)
			}

			if got, want := val, tc.val; got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
