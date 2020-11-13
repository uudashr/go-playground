package lang_test

import "testing"

func TestDeferScope_func(t *testing.T) {
	initVal, lastVal := 10, 12

	x := initVal
	defer assertEqual(t, x, initVal)

	x = lastVal
}

func TestDeferScope_anonFunc(t *testing.T) {
	initVal, lastVal := 10, 12

	x := initVal
	defer func(v int) {
		if got, want := v, initVal; got != want {
			t.Fatalf("v got: %d, want: %d", got, want)
		}
	}(x)

	x = lastVal
}

func TestDeferScope_closure(t *testing.T) {
	initVal, lastVal := 10, 12

	x := initVal
	defer func() {
		if got, want := x, lastVal; got != want {
			t.Fatalf("x got: %d, want: %d", got, want)
		}
	}()

	x = lastVal
}

func assertEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}
}
