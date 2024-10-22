package greeting_test

import (
	"testing"

	. "github.com/ovechkin-dm/mockio/mock"
	"github.com/ovechkin-dm/mockio/mockopts"
	"github.com/uudashr/go-playground/greeting"
)

func TestGreet(t *testing.T) {
	SetUp(t, mockopts.StrictVerify())
	m := Mock[greeting.Greeter]()
	WhenSingle(m.Greet("John")).ThenReturn("Hello, John!")

	msg := greeting.Greet(m, "John")
	if got, want := msg, "Hello, John!"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
