package spec_test

import (
	"fmt"
	"testing"

	"github.com/uudashr/go-playground/spec"
)

func ExampleSpec() {
	u := &User{
		Age: 20,
	}

	adultAge := AgeGreaterThan(18)
	if adultAge.SatisfiedBy(u) {
		fmt.Println("User is adult")
	}

	// Output: User is adult
}

func TestSpec(t *testing.T) {

	boolTrue := spec.OfFunc(func(v any) bool {
		return true
	})

	boolFalse := spec.OfFunc(func(v any) bool {
		return false
	})

	val := 1
	if got, want := boolTrue.SatisfiedBy(val), true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	if got, want := boolFalse.SatisfiedBy(val), false; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	notTrue := spec.Not(boolTrue)
	if got, want := notTrue.SatisfiedBy(val), false; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	notFalse := spec.Not(boolFalse)
	if got, want := notFalse.SatisfiedBy(val), true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	andTrueTrue := spec.And(boolTrue, boolTrue)
	if got, want := andTrueTrue.SatisfiedBy(val), true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	andTrueFalse := spec.And(boolTrue, boolFalse)
	if got, want := andTrueFalse.SatisfiedBy(val), false; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	andFalseFalse := spec.And(boolFalse, boolFalse)
	if got, want := andFalseFalse.SatisfiedBy(val), false; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	orTrueTrue := spec.Or(boolTrue, boolTrue)
	if got, want := orTrueTrue.SatisfiedBy(val), true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	orTrueFalse := spec.Or(boolTrue, boolFalse)
	if got, want := orTrueFalse.SatisfiedBy(val), true; got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	orFalseFalse := spec.Or(boolFalse, boolFalse)
	if got, want := orFalseFalse.SatisfiedBy(val), false; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

type User struct {
	Age int
}

func AgeGreaterThan(minAge int) spec.Spec[*User] {
	return spec.Func[*User](func(u *User) bool {
		return u.Age > minAge
	})
}
