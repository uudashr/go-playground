package main

import (
	"fmt"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

func main() {
	parsedLangs, _, err := language.ParseAcceptLanguage("ar-AE, en-US;q=0.9, fr-FR;q=0.8, id-ID;q=0.7")
	if err != nil {
		panic(err)
	}

	// for i, tag := range parsedLangs {
	// 	fmt.Printf("Tag %v, q %f\n", tag, q[i])
	// }

	langs := []language.Tag{
		parsedLangs[0],
		language.English,
		language.Arabic,
		language.French,
		language.Indonesian,
	}

	for _, langTag := range langs {
		fmt.Printf("Language: %v, Name: %v\n", langTag, langTag.String())

		unit, confidence := currency.FromTag(langTag)
		fmt.Printf("Currency: %v, confidence: %v\n", unit, confidence)

		val := 253986.2345
		unitAmount := unit.Amount(val)
		fmt.Printf("Unit Amount: %s\n", unitAmount)

		// Narrow symbol format
		narrowSymbolFormat := currency.NarrowSymbol.Default(unit)
		fmt.Printf("Narrow Symbol Format: %v\n", narrowSymbolFormat(val))

		narrowSymbolFormatStandard := narrowSymbolFormat.Kind(currency.Standard)
		fmt.Printf("Narrow Symbol Format Standard: %v\n", narrowSymbolFormatStandard(val))

		narrowSymbolFormatCash := narrowSymbolFormat.Kind(currency.Cash)
		fmt.Printf("Narrow Symbol Format Cash: %v\n", narrowSymbolFormatCash(val))

		narrowSymbolFormatAccounting := narrowSymbolFormat.Kind(currency.Accounting)
		fmt.Printf("Narrow Symbol Format Accounting: %v\n", narrowSymbolFormatAccounting(val))

		// Symbol format
		symbolFormat := currency.Symbol.Default(unit)
		fmt.Printf("Symbol Format: %v\n", symbolFormat(val))

		symbolFormatStandard := symbolFormat.Kind(currency.Standard)
		fmt.Printf("Symbol Format Standard: %v\n", symbolFormatStandard(val))

		symbolFormatCash := symbolFormat.Kind(currency.Cash)
		fmt.Printf("Symbol Format Cash: %v\n", symbolFormatCash(val))

		symbolFormatAccounting := symbolFormat.Kind(currency.Accounting)
		fmt.Printf("Symbol Format Accounting: %v\n", symbolFormatAccounting(val))

		// ISO format
		isoFormat := currency.ISO.Default(unit)
		fmt.Printf("ISO Format: %v\n", isoFormat(val))

		isoFormatStandard := isoFormat.Kind(currency.Standard)
		fmt.Printf("ISO Format Standard: %v\n", isoFormatStandard(val))

		isoFormatCash := isoFormat.Kind(currency.Cash)
		fmt.Printf("ISO Format Cash: %v\n", isoFormatCash(val))

		isoFormatAccounting := isoFormat.Kind(currency.Accounting)
		fmt.Printf("ISO Format Accounting: %v\n", isoFormatAccounting(val))

		fmt.Println("----------------")
	}

}
