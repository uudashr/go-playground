package main

import (
	"fmt"

	"github.com/cbroglie/mustache"
)

func main() {
	ctx := map[string]any{"c": "world"}
	out, err := mustache.Render("hello {{c}}", ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
