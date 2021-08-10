package contextscope_test

import (
	"context"
	"testing"
)

func TestScope(t *testing.T) {
	k1 := &scope{"1"}
	k2 := &scope{"2"}

	ctx := context.Background()
	ctx = context.WithValue(ctx, k1, "Hello")
	ctx = context.WithValue(ctx, k2, "World")

	// v1
	v1, ok := ctx.Value(k1).(string)
	if !ok {
		t.Fatal("expect value for k1")
	}

	if got, want := v1, "Hello"; got != want {
		t.Errorf("v1 got: %q, want: %q", got, want)
	}

	// v2
	v2, ok := ctx.Value(k2).(string)
	if !ok {
		t.Fatal("expect value for k2")
	}
	if got, want := v2, "World"; got != want {
		t.Errorf("v2 got: %q, want: %q", got, want)
	}
}

func TestScope_sameStructKey(t *testing.T) {
	k1 := &scope{}
	k2 := &scope{}

	ctx := context.Background()
	ctx = context.WithValue(ctx, k1, "Hello")
	ctx = context.WithValue(ctx, k2, "World")

	// v1
	v1, ok := ctx.Value(k1).(string)
	if !ok {
		t.Fatal("expect value for k1")
	}

	if got, want := v1, "Hello"; got != want {
		t.Errorf("v1 got: %q, want: %q", got, want)
	}

	// v2
	v2, ok := ctx.Value(k2).(string)
	if !ok {
		t.Fatal("expect value for k2")
	}
	if got, want := v2, "World"; got != want {
		t.Errorf("v2 got: %q, want: %q", got, want)
	}
}

type scope struct {
	id string
}
