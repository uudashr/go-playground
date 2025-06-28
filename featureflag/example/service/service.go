package service

import (
	"fmt"

	"github.com/uudashr/go-playground/featureflag"
)

type Service struct {
	name         string
	FlagOverride featureflag.BoolEvaluator
}

func New(name string) *Service {
	return &Service{name: name}
}

func (svc *Service) DoSomething() error {
	fmt.Println("Doing something in service:", svc.name)
	disabled := featureflag.BoolWithOverride("disabled", false, svc.FlagOverride)
	if disabled {
		return fmt.Errorf("service %s is disabled", svc.name)
	}

	return nil
}
