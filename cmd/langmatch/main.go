package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

// ref: https://go.dev/blog/matchlang

func main() {
	acceptLanguageHeader := "id-ID, en-US;q=0.9, ar-EG;q=0.8"
	fmt.Println("Accept-Language")
	fmt.Println("----------------")
	fmt.Println(acceptLanguageHeader)

	tags, q, err := language.ParseAcceptLanguage(acceptLanguageHeader)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("Parsed accept language")
	fmt.Println("----------------------")
	for i, tag := range tags {
		fmt.Printf("Tag %v, q %f\n", tag, q[i])
	}

	supportedLangs := []language.Tag{
		language.English,
		language.Arabic,
		language.Indonesian,
	}

	fmt.Println()
	fmt.Println("Supported languages")
	fmt.Println("-------------------")
	en := display.English.Tags()
	for _, tag := range supportedLangs {
		fmt.Printf("%-20s (%s)\n", en.Name(tag), display.Self.Name(tag))
	}

	fmt.Println()
	fmt.Println("Matched language")
	fmt.Println("----------------")
	matcher := language.NewMatcher(supportedLangs)
	tag, index, confidence := matcher.Match(tags...)
	fmt.Printf("Matched tag: %v, index: %d, confidence: %v\n", tag, index, confidence)
}
