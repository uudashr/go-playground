package main

import (
	"fmt"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func main() {
	// Promotion discount and fixed amount
	discount := 0.257457
	fixed := 20.3456

	// Print in en-US locale and language
	// Output: 26% discount applied
	// Output: USD 20.35 fixed amount
	usLang := language.MustParse("en-US")
	usPrinter := message.NewPrinter(usLang)
	usPrinter.Printf("Discount %v applied\n", number.Percent(discount))

	usUnit := currency.MustParseISO("USD")
	formatEn := currency.ISO.Default(usUnit).Kind(currency.Accounting)
	usPrinter.Printf("%v fixed amount\n", formatEn(fixed))

	fmt.Println("----------------")

	// Print in id-ID locale and language
	// Output: Diskon 25% diterapkan
	// Output: IDR 20 jumlah tetap
	idLang := language.MustParse("id-ID")
	idPrinter := message.NewPrinter(idLang)
	idPrinter.Printf("Diskon %v diterapkan\n", number.Percent(discount))

	idUnit := currency.MustParseISO("IDR")
	formatId := currency.ISO.Default(idUnit).Kind(currency.Accounting)
	idPrinter.Printf("%v jumlah tetap\n", formatId(fixed))

	fmt.Println("----------------")

	// Print in ar-AE locale and language
	// Output: تم تطبيق خصم 26%
	// Output: AED 20.35 المبلغ الثابت
	aeLang := language.MustParse("ar-AE")
	aePrinter := message.NewPrinter(aeLang)
	aePrinter.Printf("تم تطبيق خصم %v\n", number.Percent(discount))

	aeUnit := currency.MustParseISO("AED")
	formatAr := currency.ISO.Default(aeUnit).Kind(currency.Accounting)
	aePrinter.Printf("%v المبلغ الثابت\n", formatAr(fixed))

	fmt.Println("----------------")

	// Print in ar-JO locale and language
	// Output: تم تطبيق خصم 26%
	// Output: JOD 20.35 المبلغ الثابت
	joLang := language.MustParse("ar-JO")
	joPrinter := message.NewPrinter(joLang)
	joPrinter.Printf("تم تطبيق خصم %v\n", number.Percent(discount))

	joUnit := currency.MustParseISO("JOD")
	formatJo := currency.ISO.Default(joUnit).Kind(currency.Accounting)
	joPrinter.Printf("%v المبلغ الثابت\n", formatJo(fixed))
}
