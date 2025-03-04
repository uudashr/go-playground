package main

import (
	"fmt"

	"github.com/hamba/avro/v2"
)

type User struct {
	ID    int     `avro:"id"`
	Name  string  `avro:"name"`
	Email *string `avro:"email"`
}

const schemaJSON = `
{
	"type": "record",
	"name": "User",
	"fields": [
		{ "name": "id", "type": "int" },
		{ "name": "name", "type": "string" },
		{ "name": "email", "type": ["null", "string"], "default": null }
	]
}
`

func deref(s *string) string {
	if s == nil {
		return "nil"
	}

	return *s
}

func main() {
	schema, err := avro.Parse(schemaJSON)
	if err != nil {
		panic(err)
	}

	email := "alice@examle.com"
	user := User{
		ID:    1,
		Name:  "Alice",
		Email: &email,
	}

	avroData, err := avro.Marshal(schema, user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Serialized Avro Data:", avroData)

	var decodedUser User
	err = avro.Unmarshal(schema, avroData, &decodedUser)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decoded Struct: %+v\n", decodedUser)

}
