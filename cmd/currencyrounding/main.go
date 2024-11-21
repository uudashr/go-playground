package main

import (
	"fmt"

	"golang.org/x/text/currency"
)

func main() {
	val := 257874.54545

	fmt.Println("Val:", val)
	fmt.Println("=======================")
	currencyCodes := []string{"USD", "AED", "JOD", "BHD", "EGP", "PKR", "IDR", "THB", "EUR"}
	for _, code := range currencyCodes {
		unit := currency.MustParseISO(code)
		printVal(unit, val)
	}

}

func printVal(u currency.Unit, val float64) {
	fmt.Println("Amount:", u.Amount(val))

	formatISO := currency.ISO.Default(u)
	fmt.Println("ISO:", formatISO(val))

	formatISOStandard := currency.ISO.Default(u).Kind(currency.Standard)
	fmt.Println("ISO Standard:", formatISOStandard(val))

	formatISOAccounting := currency.ISO.Default(u).Kind(currency.Accounting)
	fmt.Println("ISO Accounting:", formatISOAccounting(val))

	formatISOCash := currency.ISO.Default(u).Kind(currency.Cash)
	fmt.Println("ISO Cash:", formatISOCash(val))

	fmt.Println("-----------------------")
}
