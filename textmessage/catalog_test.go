package textmessage_test

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

func ExampleSelectf() {
	cat := catalog.NewBuilder()
	cat.Set(language.English, "You are %d minutes(s) late.",
		plural.Selectf(1, "",
			plural.One, "You are 1 minute late.",
			plural.Other, "You are %d minutes late."))

	p := message.NewPrinter(language.English, message.Catalog(cat))

	p.Printf("You are %d minutes(s) late.", 1)
	p.Println()

	p.Printf("You are %d minutes(s) late.", 2)
	p.Println()

	// Output:
	// You are 1 minute late.
	// You are 2 minutes late.
}

func ExampleVar() {
	cat := catalog.NewBuilder()
	cat.Set(language.English, "You are %d minute(s) late.",
		catalog.Var("minutes",
			plural.Selectf(1, "", plural.One, "minute", plural.Other, "minutes")),
		catalog.String("You are %[1]d ${minutes} late."))

	p := message.NewPrinter(language.English, message.Catalog(cat))

	p.Printf("You are %d minute(s) late.", 1)
	p.Println()

	p.Printf("You are %d minute(s) late.", 2)
	p.Println()

	// Output:
	// You are 1 minute late.
	// You are 2 minutes late.
}

func ExampleVar_noMatch() {
	cat := catalog.NewBuilder()
	cat.Set(language.English, "You are %d minute(s) late.",
		catalog.Var("minutes", plural.Selectf(1, "", plural.One, "minute")),
		catalog.String("You are %d ${minutes} late."))

	p := message.NewPrinter(language.English, message.Catalog(cat))

	p.Printf("You are %d minute(s) late.", 1)
	p.Println()

	p.Printf("You are %d minute(s) late.", 2)
	p.Println()

	// Output:
	// You are 1 minute late.
	// You are 2  late.
}
