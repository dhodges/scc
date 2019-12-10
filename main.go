package main

import (
	"fmt"
)

func main() {
	rules, err := loadPrices()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(rules.AdPrices)
}
