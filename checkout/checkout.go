package checkout

import "github.com/dhodges/scc/pricing"

// Checkout simple checkout shopping basket
type Checkout struct {
	pricer pricing.CustomerAdPricer
	items  []pricing.AdName
}

// New create a new Checkout
func New(pricer pricing.CustomerAdPricer) Checkout {
	return Checkout{pricer: pricer}
}

// Items return a copy of items
func (c Checkout) Items() []pricing.AdName {
	items := make([]pricing.AdName, len(c.items))
	copy(items, c.items)
	return items
}

// Add add a new ad :)
func (c *Checkout) Add(ad pricing.AdName) {
	c.items = append(c.items, ad)
}

// Total return total price of all items
func (c Checkout) Total() float64 {
	return c.pricer.PriceForItems(c.items)
}
