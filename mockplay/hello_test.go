package mockplay_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/uudashr/go-playground/mockplay"
	"github.com/uudashr/go-playground/mockplay/mocks"
)

func TestHello(t *testing.T) {
	greeterMock := new(mocks.Greeter)
	greeterMock.On("Greet", "Hello").Return("Hello").Once()

	out := mockplay.Hello(greeterMock)

	if got, want := out, "Hello"; got != want {
		t.Errorf("out got %q, want %q", got, want)
	}

	mock.AssertExpectationsForObjects(t,
		greeterMock)
}
