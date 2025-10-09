package main

import (
	"fmt"
	"os"

	"golang.org/x/text/language"
)

func main() {
	acceptLang := "en_US"
	tag, q, err := language.ParseAcceptLanguage(acceptLang)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse accept language fail: %v", err)
		os.Exit(-1)
	}

	for i, t := range tag {
		base, confidence := t.Base()
		fmt.Printf("tag: %s, base_tag: %s, base_tag_confidence: %s, q: %f", t.String(), base, confidence, q[i])
	}
}
