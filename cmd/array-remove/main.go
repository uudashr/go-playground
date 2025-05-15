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

	// Observation:
	// - the options length is remain the same, but the values are changed
	// - filtered and out has smaller length, but it has same capacity just like options
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
	fmt.Println("opts:", opts, len(opts), cap(opts))
	fmt.Println("out:", out, len(out), cap(out))
	return out
}
