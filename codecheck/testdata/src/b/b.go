package main

import (
	"fmt"
)

type Account struct {
	Name        string
	Email       string
	Permissions []Permission
	Deactivated bool
}

type Permission struct {
	Domain string
	Access string
}

func main() {
	acc := Account{ // want "found 4 non keyed on struct literal \\(> 2\\)"
		"Nuruddin Ashr",
		"uudashr@gmail.com",
		[]Permission{
			Permission{"account", "read"},
			Permission{"account", "write"},
		},
		false,
	}
	fmt.Printf("%+v", acc)
}
