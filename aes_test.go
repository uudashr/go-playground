package playground_test

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"testing"
)

// https://crypto.stackexchange.com/questions/33812/is-it-acceptable-to-write-the-nonce-to-the-encrypted-file-during-aes-256-gcm

func TestEncrypt(t *testing.T) {
	strings := []string{
		"Hello World",
		"exampleplaintext",
		"andika",
		"3574434",
	}

	// Key should be 16 bytes (AES-128) or 32 (AES-256)
	// key := make([]byte, 32)
	// if _, err := io.ReadFull(rand.Reader, key); err != nil {
	// 	t.Fatal("Fail to create key:", err)
	// }

	// or use static key
	key, err := hex.DecodeString("7319b717cb7a7b2c8d36b2387b050974089637d2ac949eed94da0b20ec86cb20")
	if err != nil {
		t.Fatal("Fail to decode key:", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal("Fail to create chiper:", err)
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		t.Fatal("Fail to create nonce:", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		t.Fatal("Fail to create GCM:", err)
	}

	for _, s := range strings {
		plainText := []byte(s)

		cipherText := aesgcm.Seal(nil, nonce, plainText, nil)
		encodedCipherText := hex.EncodeToString(cipherText)
		t.Logf("Cipher text len=%d %q\n", len(cipherText), encodedCipherText)

		// nonceSize := aesgcm.NonceSize()
		// decipherNonce, byteText := cipherText[:nonceSize], cipherText[nonceSize:]
		// dechiperText, err := aesgcm.Open(nil, decipherNonce, byteText, nil)
		// if err != nil {
		// 	t.Errorf("Fail to decipher %q\n", plainText)
		// 	continue
		// }

		dechiperText, err := aesgcm.Open(nil, nonce, cipherText, nil)
		if err != nil {
			t.Errorf("Fail to decipher %q: %v\n", plainText, err)
		}

		t.Logf("Dechiper text %q", dechiperText)
	}

}

func TestAesDecrypt(t *testing.T) {
	cipherTextEncoded := "C4HPsx3oUbRvgCJ2Q4Hb0g3dSPRJ0za++n2A"
	nonceEncoded := "weGdr6jwx8bckyVs"

	key, err := base64.StdEncoding.DecodeString("WuVnt3tvZGOgo70OJvbbVmOcSD7yLiDYuBioeDCt5jY=")
	if err != nil {
		t.Fatal("Fail to decode key:", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal("Fail to create chiper:", err)
	}

	nonce, err := base64.StdEncoding.DecodeString(nonceEncoded)
	if err != nil {
		t.Fatal("Fail to decode nonce:", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		t.Fatal("Fail to create GCM:", err)
	}

	cipherText, err := base64.StdEncoding.DecodeString(cipherTextEncoded)
	if err != nil {
		t.Fatal("Fail to decode cipherText:", err)
	}

	dechiperText, err := aesgcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		t.Errorf("Fail to decipher %q: %v\n", cipherTextEncoded, err)
	}

	t.Logf("Dechiper text %q", dechiperText)
}
