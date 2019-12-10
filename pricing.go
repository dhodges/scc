package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// AdName stored as a string
type AdName string

// AdGroupPrice ads priced by grouping, e.g. 3 for 2
type AdGroupPrice struct {
	GroupCount int
	Multiplier int
}

// Customer stored as a string
type Customer string

// CustomerPricingRules Discount and Group ad prices for a given Customer
type CustomerPricingRules struct {
	DiscountPrices map[AdName]float64
	GroupPrices    map[AdName]AdGroupPrice
}

// PricingRules contains all known ad prices and special customer deals
type PricingRules struct {
	AdPrices map[AdName]float64
	Deals    map[Customer]CustomerPricingRules
}

// PriceRules holds all known ad prices
var PriceRules PricingRules

// LoadPrices gather all known ad prices, including customer deals
func loadPrices() (PricingRules, error) {
	data, err := ioutil.ReadFile("./pricing_rules.json")
	if err != nil {
		fmt.Print(err)
	} else {
		err = json.Unmarshal(data, &PriceRules)
		if err != nil {
			fmt.Print(err)
		}
	}
	return PriceRules, err
}
