package service_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uudashr/go-playground/featureflag"
	"github.com/uudashr/go-playground/featureflag/example/service"
)

func TestService_DoSomething(t *testing.T) {
	tests := []struct {
		name         string
		flagOverride featureflag.BoolEvaluator
		assertError  require.ErrorAssertionFunc
	}{
		{
			name:        "succeed",
			assertError: require.NoError,
		},
		{
			name: "disabled service",
			flagOverride: &mapFlags{map[string]any{
				"disabled": true,
			}},
			assertError: require.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			svc := service.New("hello")

			// When we decide not to experiment, it require us to remove this as well.
			// It is effort, but not so much effort to remove this line.
			// Having this svc.FlagOverrid will not cause any race-condition/parallel test issue.
			svc.FlagOverride = tt.flagOverride

			// Act
			err := svc.DoSomething()

			// Assert
			tt.assertError(t, err)
		})
	}
}

type mapFlags struct {
	flags map[string]any
}

func (m *mapFlags) Bool(name string, defaultValue bool) bool {
	if val, ok := m.flags[name]; ok {
		return val.(bool)
	}

	return defaultValue
}

func (m *mapFlags) String(name string, defaultValue string) string {
	if val, ok := m.flags[name]; ok {
		return val.(string)
	}

	return defaultValue
}
