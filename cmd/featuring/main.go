package main

import (
	"context"
	"fmt"

	"github.com/open-feature/go-sdk/openfeature"
	"github.com/open-feature/go-sdk/openfeature/memprovider"
)

func main() {
	provider := memprovider.NewInMemoryProvider(map[string]memprovider.InMemoryFlag{
		"v2_enabled": {
			Key:            "v2_enabled",
			State:          memprovider.Enabled,
			DefaultVariant: "false",
			Variants: map[string]any{
				"true":  true,
				"false": false,
			},
		},
	})
	openfeature.SetProviderAndWait(provider)
	defer openfeature.Shutdown()

	client := openfeature.NewClient("app")
	v2Enabled := client.Boolean(
		context.Background(), "v2_enabled", true, openfeature.EvaluationContext{},
	)

	fmt.Println("v2_enabled:", v2Enabled)
}
