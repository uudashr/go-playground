package jwebtoken

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type SecureClaim struct {
	jwt.StandardClaims
	UserFingerprint string `json:"aud,omitempty"`
}

func Login(issuer string, w http.ResponseWriter) error {
	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}

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
	secretKey := []byte("secret")
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"accessToken": signedToken,
	})
}

func Validate(r *http.Request) error {
	cookie, err := r.Cookie("__Secure-Fgp")
	if err != nil {
		return err
	}

	tokenString := r.Header.Get("Access-Token")
	if tokenString == "" {
		return errors.New("no token found")
	}

	userFingerprint := cookie.Value

	digest := sha256.New()
	userFingerprintDigest := digest.Sum([]byte(userFingerprint))
	userFingerprintHash := hex.EncodeToString(userFingerprintDigest)

	var claim SecureClaim
	_, err = jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})
	if err != nil {
		return err
	}

	if claim.UserFingerprint == userFingerprintHash {
		return errors.New("invalid token")
	}

	return claim.Valid()
}
