package main

import (
	"fmt"

	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

func main() {
	runNumbers()
}

func runNumbers() {
	for _, lang := range []string{"en", "de", "de-CH", "fr", "bn"} {
		p := message.NewPrinter(language.Make(lang))
		p.Printf("%-6s %g\n", lang, 123456.78)
	}
}

func runSubstitutions() {
	cat := catalog.NewBuilder()
	cat.SetString(language.Dutch, "You have chosen to play %m.", "U heeft ervoor gekozen om %m te spelen.")
	cat.SetString(language.Dutch, "basketball", "basketbal")
	cat.SetString(language.Dutch, "hockey", "ijshockey")
	cat.SetString(language.Dutch, "soccer", "voetbal")
	cat.SetString(language.BritishEnglish, "soccer", "football")

	for _, sport := range []string{"soccer", "basketball", "hockey", "baseball"} {
		for _, lang := range []string{"en", "en-GB", "nl"} {
			p := message.NewPrinter(language.Make(lang), message.Catalog(cat))
			fmt.Printf("%-6s %s\n", lang, p.Sprintf("You have chosen to play %m.", sport))
		}
		fmt.Println()
	}
}

func runSimple() {
	cat := catalog.NewBuilder()
	cat.Set(language.English, "You are %d minute(s) late.",
		plural.Selectf(1, "",
			plural.One, "You are 1 minute late.",
			plural.Other, "You are %d minutes late."))

	p := message.NewPrinter(language.English, message.Catalog(cat))
	p.Printf("You are %d minute(s) late.\n", 1)
}
