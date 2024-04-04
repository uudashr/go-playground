package password_test

import (
	"testing"

	"github.com/uudashr/go-playground/password"
)

func TestHashPassword(t *testing.T) {
	salt, err := password.GenerateSalt(16)
	if err != nil {
		t.Fatal(salt)
	}

	hash := password.Hash("secret", salt)

	if !password.CheckPassword("secret", salt, hash) {
		t.Error("password not match")
	}
}
