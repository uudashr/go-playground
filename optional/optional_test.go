package optional_test

import (
	"testing"

	"github.com/uudashr/go-playground/optional"
)

func TestOptional(t *testing.T) {
	s := "Hello"
	opt := optional.Ok(s)
	value, ok := opt.Get()
	if !ok {
		t.Fatalf("expected ok")
	}

	if got, want := value, s; got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if got, want := opt.Empty(), false; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := opt.Default("Holla"), "Hello"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
