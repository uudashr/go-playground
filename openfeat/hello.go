package openfeat

import (
	"context"

	"github.com/open-feature/go-sdk/openfeature"
)

func Hello() string {
	client := openfeature.NewClient("hello")
	msg := client.String(context.Background(), "message", "hello", openfeature.EvaluationContext{})
	return msg
}
