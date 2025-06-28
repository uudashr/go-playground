package openfeat

import (
	"testing"

	"github.com/open-feature/go-sdk/openfeature"
	"github.com/open-feature/go-sdk/openfeature/memprovider"
	feattesting "github.com/open-feature/go-sdk/openfeature/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	testProvider := feattesting.NewTestProvider()
	err := openfeature.SetProviderAndWait(testProvider)
	require.NoError(t, err)

	tests := []struct {
		name          string
		flags         map[string]memprovider.InMemoryFlag
		assertError   require.ErrorAssertionFunc
		expectMessage string
	}{
		{
			name:          "normal",
			flags:         map[string]memprovider.InMemoryFlag{},
			expectMessage: "hello",
		},
		{
			name: "override by flag",
			flags: map[string]memprovider.InMemoryFlag{
				"message": memprovider.InMemoryFlag{
					DefaultVariant: "default",
					Variants: map[string]any{
						"default": "hi",
					},
				},
			},
			expectMessage: "hi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testProvider.UsingFlags(t, tt.flags)
			defer testProvider.Cleanup()

			msg := Hello()
			assert.Equal(t, tt.expectMessage, msg)
		})
	}
}
