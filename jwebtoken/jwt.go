package jwebtoken

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// ref: https://hasura.io/blog/best-practices-of-using-jwt-with-graphql
//      https://cheatsheetseries.owasp.org/cheatsheets/JSON_Web_Token_for_Java_Cheat_Sheet.html

type SecureClaim struct {
	jwt.StandardClaims
	UserFingerprint string `json:"userFingerprint,omitempty"`
}

func Authorize(issuer, secret string, w http.ResponseWriter) error {
	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}

	// Create user fingerprint as extra security
	userFingerprint := hex.EncodeToString(b)
	http.SetCookie(w, &http.Cookie{
		Name:     "__Secure-Fgp",
		Value:    userFingerprint,
		SameSite: http.SameSiteStrictMode,
	})

	digest := sha256.New()
	userFingerprintDigest := digest.Sum([]byte(userFingerprint))
	userFingerprintHash := hex.EncodeToString(userFingerprintDigest)

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, SecureClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: now.Add(15 * time.Minute).Unix(),
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
		},
		UserFingerprint: userFingerprintHash,
	})
	secretKey := []byte(secret)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"accessToken": signedToken,
	})
}

func Validate(secret string, r *http.Request) error {
	// Get access token
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return errors.New("no authorization header")
	}

	authSegments := strings.SplitN(authHeader, " ", 2)
	if len(authSegments) < 2 {
		return errors.New("invalid authorization header")
	}

	if authSegments[0] != "Bearer" {
		return errors.New("invalid authentication scheme")
	}

	tokenString := authSegments[1]

	// Get user fingerprint
	cookie, err := r.Cookie("__Secure-Fgp")
	if err != nil {
		return err
	}

	userFingerprint := cookie.Value

	digest := sha256.New()
	userFingerprintDigest := digest.Sum([]byte(userFingerprint))
	userFingerprintHash := hex.EncodeToString(userFingerprintDigest)

	// Validate both token and user fingerprint
	var claim SecureClaim
	_, err = jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return err
	}

	if claim.UserFingerprint == userFingerprintHash {
		return errors.New("invalid token")
	}

	return claim.Valid()
}
