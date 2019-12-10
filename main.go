package main

import (
	"fmt"

	"github.com/dhodges/scc/pricing"
)

func main() {
	rules, err := pricing.FetchRules()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(rules.AdPrices)
}
