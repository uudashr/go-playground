package injection_test

import (
	"fmt"
	"testing"

	"github.com/uudashr/go-playground/injection"
)

func ExampleRegistry() {
	reg := injection.NewRegistry()

	reg.ProvideVal("John", "name")
	reg.ProvideVal("Hello", "greeting")

	reg.InjectFunc(func(name, greeting string) {
		fmt.Printf("%s %s", greeting, name)
	}, "name", "greeting")

	// Output: Hello John
}

func TestResolve(t *testing.T) {
	testCases := map[string]struct {
		value string
	}{
		"simple val": {
			value: "John",
		},
		"another simple val": {
			value: "Doe",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			reg := injection.NewRegistry()
			reg.ProvideVal(tc.value, "name")
			v, err := reg.Resolve("name")
			if err != nil {
				t.Fatal("Resolve fail: ", err)
			}

			if got, want := v, tc.value; got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestGet_ProvideFunc(t *testing.T) {
	testCases := map[string]struct {
		fn    any
		value string
	}{
		"simple val": {
			fn: func() string {
				return "John"
			},
			value: "John",
		},
		"another simple val": {
			fn: func() string {
				return "Doe"
			},
			value: "Doe",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			reg := injection.NewRegistry()
			err := reg.Provide(tc.fn, "name")
			if err != nil {
				t.Fatal("Provide fail: ", err)
			}

			v, err := reg.Resolve("name")
			if err != nil {
				t.Fatal("Resolve fail: ", err)
			}

			if got, want := v, tc.value; got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestGet_ProvideFunc_args(t *testing.T) {
	type provider struct {
		val       any
		label     injection.Label
		argLabels []injection.Label
	}
	testCases := map[string]struct {
		providers []provider
		label     injection.Label
		output    any
	}{
		"simple arg": {
			providers: []provider{
				{
					val:   "John",
					label: "name",
				},
				{
					val: func(name string) string {
						return "Hello " + name
					},
					label:     "greeting",
					argLabels: []injection.Label{"name"},
				},
			},
			label:  "greeting",
			output: "Hello John",
		},
		"another simple arg": {
			providers: []provider{
				{
					val:   "Bob",
					label: "name",
				},
				{
					val: func(name string) string {
						return "Hi " + name
					},
					label:     "greeting",
					argLabels: []injection.Label{"name"},
				},
			},
			label:  "greeting",
			output: "Hi Bob",
		},
		"two args": {
			providers: []provider{
				{
					val:   "John",
					label: "name",
				},
				{
					val:   "Hello",
					label: "greeting",
				},
				{
					val: func(name, greeting string) string {
						return greeting + " " + name + ", good morning!"
					},
					label:     "morningGreet",
					argLabels: []injection.Label{"name", "greeting"},
				},
			},
			label:  "morningGreet",
			output: "Hello John, good morning!",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			reg := injection.NewRegistry()
			for _, p := range tc.providers {
				err := reg.Provide(p.val, p.label, p.argLabels...)
				if err != nil {
					t.Fatal("Provide fail: ", err)
				}
			}

			v, err := reg.Resolve(tc.label)
			if err != nil {
				t.Fatal("Resolve fail:", err)
			}

			if got, want := v, tc.output; got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestInject_Func(t *testing.T) {
	type valueLabel struct {
		label injection.Label
		value string
	}
	testCases := map[string]struct {
		values []valueLabel
	}{
		"simple injection": {
			values: []valueLabel{
				{label: "name", value: "John"},
				{label: "greeting", value: "Hello"},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			reg := injection.NewRegistry()
			for _, v := range tc.values {
				err := reg.Provide(v.value, v.label)
				if err != nil {
					t.Fatal("Provide fail: ", err)
				}
			}

			var (
				injected             bool
				gotName, gotGreeting string
			)

			labels := make([]injection.Label, len(tc.values))
			for i, v := range tc.values {
				labels[i] = v.label
			}

			err := reg.InjectFunc(func(name, greeting string) {
				gotName = name
				gotGreeting = greeting
				injected = true
			}, labels...)
			if err != nil {
				t.Fatal("Inject fail: ", err)
			}

			if !injected {
				t.Fatal("injection not called")
			}

			if got, want := gotName, tc.values[0].value; got != want {
				t.Errorf("%q got %v, want %v", tc.values[0].label, got, want)
			}

			if got, want := gotGreeting, tc.values[1].value; got != want {
				t.Errorf("%q got %q, want %q", tc.values[0].label, got, want)
			}
		})
	}
}

func TestInject_Struct(t *testing.T) {
	type User struct {
		Name       string `injection:"name"`
		Salutation string `injection:"salutation"`
	}

	reg := injection.NewRegistry()

	reg.ProvideVal("John", "name")
	reg.ProvideVal("Mr.", "salutation")

	var user User
	err := reg.Inject(&user)
	if err != nil {
		t.Fatal("Inject fail:", err)
	}

	if got, want := user.Name, "John"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if got, want := user.Salutation, "Mr."; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestProvide_Type(t *testing.T) {
	reg := injection.NewRegistry()

	hg := &HiGreeter{}
	reg.ProvideVal(hg, "greeter")

	var greetResult string
	reg.InjectFunc(func(greeter *HiGreeter) {
		greetResult = greeter.Greet("John")
	}, "greeter")

	if got, want := greetResult, "Hi John"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

type HiGreeter struct {
}

func (hg *HiGreeter) Greet(name string) string {
	return "Hi " + name
}
