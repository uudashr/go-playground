package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

// ref: https://pkg.go.dev/golang.org/x/text

func main() {
	discount := 0.257
	fmt.Println("Discount:", discount)

	msg := &Message{}

	// Print in en-US locale and language
	// Output: 25.7% discount applied
	msg.DiscountApplied(discount, "en-US")

	// Print in en-GB locale and language
	// Output: 25.7% discount applied
	msg.DiscountApplied(discount, "en-GB")

	// Print in id-ID locale and language
	// Output: Diskon 25,7% diterapkan
	msg.DiscountApplied(discount, "id-ID")

	// Print in de-DE locale and language
	// Output: 25,7% Rabatt angewendet
	msg.DiscountApplied(discount, "de-DE")

	// Print in fr-FR locale and language
	// Output: 25,7% de réduction appliquée
	msg.DiscountApplied(discount, "fr-FR")

	// Print in ar-SA locale and language
	// Output: تم تطبيق خصم بنسبة ٢٥٫٧٪
	msg.DiscountApplied(discount, "ar-SA")

	// Print in ar-AE locale and language
	// Output: ٢٥٫٧٪ discount applied
	msg.DiscountApplied(discount, "ar-AE")

	// Print in en-US locale and language
	// Output: Good morning
	msg.Greeting("en-US")

	// Print in id-ID locale and language
	// Output: Selamat pagi
	msg.Greeting("id-ID")

	// Print in ar-SA locale and language
	// Output: صباح الخير
	msg.Greeting("ar-SA")

	// or use https://lokalise.com/blog/go-internationalization-using-go-i18n/
}

type Message struct {
}

func (m *Message) DiscountApplied(discount float64, locale string) {
	var p *message.Printer
	switch locale {
	case "en-US":
		p = message.NewPrinter(language.AmericanEnglish)
	case "en-GB":
		p = message.NewPrinter(language.BritishEnglish)
	case "id-ID":
		p = message.NewPrinter(language.Indonesian)
	case "de-DE":
		p = message.NewPrinter(language.German)
	case "fr-FR":
		p = message.NewPrinter(language.French)
	case "ar-SA":
		p = message.NewPrinter(language.Arabic)
	case "ar-AE":
		p = message.NewPrinter(language.Arabic)
	default:
		p = message.NewPrinter(language.AmericanEnglish)
	}

	switch locale {
	case "en-US", "en-GB":
		p.Printf("%v discount applied\n", number.Percent(discount))
	case "id-ID":
		p.Printf("Diskon %v diterapkan\n", number.Percent(discount))
	case "de-DE":
		p.Printf("%v Rabatt angewendet\n", number.Percent(discount))
	case "fr-FR":
		p.Printf("%v de réduction appliquée\n", number.Percent(discount))
	case "ar-SA":
		p.Printf("تم تطبيق خصم بنسبة %v\n", number.Percent(discount))
	case "ar-AE":
		p.Printf("%v discount applied\n", number.Percent(discount))
	default:
		p.Printf("%v discount applied\n", number.Percent(discount))
	}

	// switch locale {
	// case "en-US", "en-GB":
	// 	p.Printf("%.1f%% discount applied\n", discount*100)
	// case "id-ID":
	// 	p.Printf("Diskon %.1f%% diterapkan\n", discount*100)
	// case "de-DE":
	// 	p.Printf("%.1f%% Rabatt angewendet\n", discount*100)
	// case "fr-FR":
	// 	p.Printf("%.1f%% de réduction appliquée\n", discount*100)
	// case "ar-SA":
	// 	p.Printf("تم تطبيق خصم بنسبة %.1f%%\n", discount*100)
	// default:
	// 	p.Printf("%.1f%% discount applied\n", discount*100)
	// }
}

func (m *Message) Greeting(locale string) {
	switch locale {
	case "en-US":
		fmt.Println("Good morning")
	case "id-ID":
		fmt.Println("Selamat pagi")
	case "ar-SA":
		fmt.Println("صباح الخير")
	default:
		fmt.Println("Good morning")
	}
}
