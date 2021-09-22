package jwebtoken_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestSigning(t *testing.T) {
	// or use jwt.StandardClaims instead of jwt.MapClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "1234567890",
		"name": "John Doe",
		"iss":  "Jungle",
		"aud":  "jungleapp",
		"iat":  time.Now().Unix(),
	})
	secretKey := []byte("secret")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		t.Fatal("Fail to sign:", err)
	}

	// or use jwt.ParseWithClaims() if want to decode to supplied claims struct
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		t.Fatal("Parse fail:", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatal("Invalid claim type")
	}

	if !parsedToken.Valid {
		t.Fatal("Invalid token")
	}

	if got, want := claims["sub"], "1234567890"; got != want {
		t.Errorf("Claim sub got: %q, want: %q", got, want)
	}
}
