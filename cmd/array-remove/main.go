package main

import (
	"fmt"
)

func main() {
	options := []string{"credit-card", "cash", "wallet"}
	fmt.Println("--- Before ---")
	fmt.Printf("options: %v len=%d cap=%d\n", options, len(options), cap(options))

	fmt.Println("--- After ---")
	filtered := filterPaymentTypes(options)
	fmt.Printf("options: %v len=%d cap=%d\n", options, len(options), cap(options))
	fmt.Printf("filtered: %v len=%d cap=%d\n", filtered, len(filtered), cap(filtered))

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
	fmt.Printf("opts: %v len=%d cap=%d\n", opts, len(opts), cap(opts))
	fmt.Printf("out: %v len=%d\n", out, len(out), cap(out))
	return out
}
