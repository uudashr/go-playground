package securerandom_test

import (
	"crypto/rand"
	"encoding/base64"
	"reflect"
	"testing"
)

func TestRandom(t *testing.T) {
	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	// or use hex.EncodeToString()
	encoded := base64.URLEncoding.EncodeToString(b)

	// decode, also can use hex.DecodeString()
	decoded, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		t.Fatal("Decode fail:", err)
	}

	if got, want := decoded, b; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
