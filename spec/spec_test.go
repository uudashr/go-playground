package spec_test

import (
	"fmt"

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

type User struct {
	Age int
}

func AgeGreaterThan(minAge int) spec.Spec[*User] {
	return spec.Func[*User](func(u *User) bool {
		return u.Age > minAge
	})
}
