package serializejson_test

import (
	"bytes"
	"encoding/json"
	"testing"
)

type user struct {
	Name  string
	Email string
}

func BenchmarkMarshal(b *testing.B) {
	u := &user{
		Name:  "John Appleseed",
		Email: "john.appleseed@nowhere.com",
	}

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(u)
		if err != nil {
			b.Fatal("Marsal error:", err)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	u := &user{
		Name:  "John Appleseed",
		Email: "john.appleseed@nowhere.com",
	}

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		err := enc.Encode(u)
		if err != nil {
			b.Fatal("Marsal error:", err)
		}
	}
}
