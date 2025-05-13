package main

import (
	"fmt"
)

func main() {
	options := []string{"credit-card", "cash", "wallet"}
	fmt.Println(options)
	filtered := filterPaymentTypes(options)
	fmt.Println(options)
	fmt.Println(filtered)
}

func filterPaymentTypes(opts []string) []string {
	for i := 0; i < len(opts); i++ {
		if string(opts[i]) == "cash" {
			opts = append(opts[:i], opts[i+1:]...)
			break
		}
	}
	return opts
}
