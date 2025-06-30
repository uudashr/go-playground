package main

import (
	"context"
	"os"
)

// go get -tool github.com/a-h/templ/cmd/templ@latest
// go tool templ generate

func main() {
	component := hello("John")
	component.Render(context.Background(), os.Stdout)
}
