package main

import (
	"fmt"
)

func main() {
	options := []string{"credit-card", "cash", "wallet"}
	fmt.Println("options:", options, len(options), cap(options))
	filtered := filterPaymentTypes(options)
	fmt.Println("options:", options, len(options), cap(options))
	fmt.Println("filtered:", filtered, len(filtered), cap(filtered))
}

func filterPaymentTypes(opts []string) []string {
	var out []string
	for i := 0; i < len(opts); i++ {
		if string(opts[i]) == "cash" {
			out = append(opts[:i], opts[i+1:]...)
			// out = slices.Delete(opts, i, i+1)
			break
		}
	}
	fmt.Println("out:", out, len(out), cap(out))
	return out
}
