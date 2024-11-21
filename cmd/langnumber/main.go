package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

type values struct {
	Decimal float64
	Percent float32
}

func main() {
	vals := values{
		Decimal: 257874.54545,
		Percent: 0.25876,
	}

	languagesCodes := []string{"en", "id", "ar", "de", "es", "fr", "ja", "ko", "pt", "ru", "zh"}

	fmt.Println("Decimal:", vals.Decimal)
	fmt.Println("Percent:", vals.Percent)

	namer := display.English.Tags()
	for _, code := range languagesCodes {
		tag := language.MustParse(code)
		printVal(tag, vals, namer)
	}
}

func printVal(t language.Tag, vals values, namer display.Namer) {
	fmt.Printf("== Lang: %s (%s) ===\n", t, namer.Name(t))
	p := message.NewPrinter(t)
	p.Printf("Decimal: %v\n", number.Decimal(vals.Decimal))
	p.Printf("Percent: %v\n", number.Percent(vals.Percent))
	p.Printf("Engineering: %v\n", number.Engineering(vals.Decimal))
	p.Printf("Scientific: %v\n", number.Scientific(vals.Decimal))
}
