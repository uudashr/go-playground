package main

import (
	"fmt"

	"github.com/biter777/countries"
)

func main() {
	countryCode := countries.ByName("BH")
	fmt.Println(countryCode)
}
