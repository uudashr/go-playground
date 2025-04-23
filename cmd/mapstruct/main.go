package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-viper/mapstructure/v2"
)

func main() {
	p := Person{
		Name: "John Doe",
		Age:  30,
	}

	fmt.Printf("Struct: %+v\n", p)

	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println(string(b))

	var m map[string]any
	err = json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("'name':", reflect.TypeOf(m["name"]))
	fmt.Println("'age_in_years':", reflect.TypeOf(m["age_in_years"]))

	var dp Person

	config := &mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &dp,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		fmt.Println("Error creating decoder:", err)
		return
	}

	err = decoder.Decode(m)
	if err != nil {
		fmt.Println("Error decoding mapstructure:", err)
		return
	}

	fmt.Printf("Decoded struct: %+v\n", dp)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age_in_years"`
}
