package main

import (
	"encoding/json"
	"os"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID    string          `json:"id"`
	Price decimal.Decimal `json:"price,omitempty"`
}

func main() {
	product := Product{
		ID: "123",
		// Price: decimal.NewFromFloat(123.45),
	}

	if err := json.NewEncoder(os.Stdout).Encode(product); err != nil {
		panic(err)
	}
}
