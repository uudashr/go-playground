package password_test

import (
	"bytes"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBcrypt(t *testing.T) {
	hash1, err := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	hash2, err := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(hash1, hash2) {
		t.Errorf("hashes are the same: %v", hash1)
	}

	err = bcrypt.CompareHashAndPassword(hash1, []byte("secret"))
	if err != nil {
		t.Errorf("hash1 password does not match: %v", err)
	}

	err = bcrypt.CompareHashAndPassword(hash2, []byte("secret"))
	if err != nil {
		t.Errorf("hash2 password does not match: %v", err)
	}
}
