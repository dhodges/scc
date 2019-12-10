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

// CustomerRules for a specific Customer only
type CustomerRules struct {
	CustomerName Customer
	AdPrices     map[AdName]float64
	Deals        CustomerPricingRules
}

// CustomerAdPricer knows how to calculate customer ad prices, including special deals
type CustomerAdPricer interface {
	PriceForItems(ads []AdName) float64
}
