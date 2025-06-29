package service

import (
	"fmt"

	"github.com/uudashr/go-playground/featureflag"
)

type Service struct {
	name string

	// FlagOverride to override the feature flag evaluation.
	// Primarily used for testing purposes to simulate different flag states.
	// More likely is not going to be used in production code, but it can.
	FlagOverride featureflag.BoolEvaluator
}

func New(name string) *Service {
	return &Service{name: name}
}

func (svc *Service) DoSomething() error {
	fmt.Println("Doing something in service:", svc.name)

	// Removing flag will need to remove this line along with the FlagOverride field.
	// Not so much effort isn't it?
	disabled := featureflag.BoolWithOverride("disabled", false, svc.FlagOverride)
	if disabled {
		return fmt.Errorf("service %s is disabled", svc.name)
	}

	return nil
}
