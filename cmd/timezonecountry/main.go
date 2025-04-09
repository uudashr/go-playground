package main

import (
	"fmt"
	"time"

	"github.com/go-playground/tz"
)

func main() {
	countryCodes := []string{
		"BH",
		"AE",
		"EG",
		"SA",
		"QA",
		"JO",
		"PK",
		"ID",
	}

	for _, code := range countryCodes {
		fmt.Println("---------------------------------")
		country, found := tz.GetCountry(code)
		if !found {
			fmt.Printf("country %s not found\n", code)
			continue
		}

		fmt.Printf("Country code: %s\n", country.Code)
		fmt.Printf("Country name: %s\n", country.Name)
		fmt.Printf("Country zones: %v\n", country.Zones)
		fmt.Printf("Dump: %s\n", country)
		fmt.Printf("First zone: %s\n", country.Zones[0].Name)

		loc, err := time.LoadLocation(country.Zones[0].Name)
		if err != nil {
			fmt.Printf("Error loading location: %v\n", err)
			continue
		}
		fmt.Printf("Location: %s\n", loc)

	}
}
