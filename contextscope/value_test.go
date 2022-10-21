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

func TestScope_override(t *testing.T) {
	key := &scope{}

	ctx1 := context.Background()
	ctx2 := context.WithValue(ctx1, key, "Hello")
	ctx3 := context.WithValue(ctx2, key, "World")

	// ctx1
	_, ok := ctx1.Value(key).(string)
	if ok {
		t.Fatal("expect no value on ctx1")
	}

	// ctx2
	v2, ok := ctx2.Value(key).(string)
	if !ok {
		t.Fatal("expect value on ctx2")
	}

	if got, want := v2, "Hello"; got != want {
		t.Errorf("v2 got: %q, want: %q", got, want)
	}

	// ctx3
	v3, ok := ctx3.Value(key).(string)
	if !ok {
		t.Fatal("expect value on ctx3")
	}

	if got, want := v3, "World"; got != want {
		t.Errorf("v3 got: %q, want: %q", got, want)
	}
}

type scope struct {
	id string
}
