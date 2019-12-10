package pricing

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

// Rules contains all known ad prices and special customer deals
type Rules struct {
	AdPrices map[AdName]float64
	Deals    map[Customer]CustomerPricingRules
}

// FetchRules get the current state of all pricing rules
// for now, let's just load this from a JSON file
func FetchRules() (Rules, error) {
	return loadPricesFromJSON()
}
