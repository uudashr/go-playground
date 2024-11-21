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

	// Print in us-US locale and language
	// Output: 26% discount applied
	// Output: USD 20.35 fixed amount
	us, _ := language.Parse("en-US")
	usp := message.NewPrinter(us)
	usp.Printf("Discount %v applied\n", number.Percent(discount))

	enUnit, _ := currency.ParseISO("USD")
	formatEn := currency.ISO.Default(enUnit).Kind(currency.Accounting)
	usp.Printf("%v fixed amount\n", formatEn(fixed))

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

	// Print in ae-AE locale and language
	// Output: تم تطبيق خصم 26%
	// Output: AED 20.35 المبلغ الثابت
	ae, _ := language.Parse("ar-AE")
	aep := message.NewPrinter(ae)
	aep.Printf("تم تطبيق خصم %v\n", number.Percent(discount))

	arUnit, _ := currency.ParseISO("AED")
	formatAr := currency.ISO.Default(arUnit).Kind(currency.Accounting)
	aep.Printf("%v المبلغ الثابت\n", formatAr(fixed))

	fmt.Println("----------------")

	// Print in ar-JO locale and language
	// Output: تم تطبيق خصم 26%
	// Output: JOD 20.35 المبلغ الثابت
	jo, _ := language.Parse("ar-JO")
	jop := message.NewPrinter(jo)
	jop.Printf("تم تطبيق خصم %v\n", number.Percent(discount))

	joUnit, _ := currency.ParseISO("JOD")
	formatJo := currency.ISO.Default(joUnit).Kind(currency.Accounting)
	jop.Printf("%v المبلغ الثابت\n", formatJo(fixed))
}
