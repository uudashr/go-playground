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
	en, _ := language.Parse("en-US")
	enp := message.NewPrinter(en)
	enp.Printf("Discount %v applied\n", number.Percent(discount))

	enUnit, _ := currency.ParseISO("USD")
	formatEn := currency.ISO.Default(enUnit).Kind(currency.Accounting)
	enp.Printf("%v fixed amount\n", formatEn(fixed))

	fmt.Println("----------------")

	// Print in id-ID locale and language
	// Output: Diskon 25% diterapkan
	// Output: IDR 20 jumlah tetap
	id, _ := language.Parse("id-ID")
	idp := message.NewPrinter(id)
	idp.Printf("Diskon %v diterapkan\n", number.Percent(discount))

	idUnit, _ := currency.ParseISO("IDR")
	formatId := currency.ISO.Default(idUnit).Kind(currency.Accounting)
	idp.Printf("%v jumlah tetap\n", formatId(fixed))

	fmt.Println("----------------")

	// Print in ar-AE locale and language
	// Output: تم تطبيق خصم 26%
	// Output: AED 20.35 المبلغ الثابت
	ar, _ := language.Parse("ar-AE")
	arp := message.NewPrinter(ar)
	arp.Printf("تم تطبيق خصم %v\n", number.Percent(discount))

	arUnit, _ := currency.ParseISO("AED")
	formatAr := currency.ISO.Default(arUnit).Kind(currency.Accounting)
	arp.Printf("%v المبلغ الثابت\n", formatAr(fixed))
}
