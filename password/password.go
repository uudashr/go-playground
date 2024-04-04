package password

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/pbkdf2"
)

func GenerateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}

func Hash(password string, salt []byte) []byte {
	iterations := 10000
	keyLength := 32

	return pbkdf2.Key([]byte(password), salt, iterations, keyLength, sha256.New)
}

func HashBase64(password string, salt []byte) string {
	hash := Hash(password, salt)
	return base64.StdEncoding.EncodeToString(hash)
}

func CheckPassword(password string, salt []byte, hashedPassword []byte) bool {
	return bytes.Equal(Hash(password, salt), hashedPassword)
}

func CheckPasswordBase64Hash(password string, salt []byte, hashedPassword string) bool {
	return HashBase64(password, salt) == hashedPassword
}
