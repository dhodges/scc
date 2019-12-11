package checkout

import (
	"reflect"
	"testing"

	"github.com/dhodges/scc/pricing"
)

func getCustomerPriceRules(t *testing.T, customer pricing.Customer) pricing.CustomerAdPricer {
	rules, err := pricing.FetchPricingRules(customer)
	if err != nil {
		t.Errorf("error when fetching customer ad price rules: %v", err)
	}
	return rules
}

func TestAddItems(t *testing.T) {
	customer := pricing.Customer("SecondBite")
	pricingRules := getCustomerPriceRules(t, customer)
	ads := []pricing.AdName{
		pricing.AdName("classic"),
		pricing.AdName("standout"),
		pricing.AdName("premium"),
	}
	co := New(pricingRules)
	for _, item := range ads {
		co.Add(item)
	}

	if !reflect.DeepEqual(co.Items(), ads) {
		t.Errorf("checkout items are wrong")
	}
}

func TestTotal(t *testing.T) {
	customer := pricing.Customer("SecondBite")
	pricingRules := getCustomerPriceRules(t, customer)
	ads := []pricing.AdName{
		pricing.AdName("classic"),
		pricing.AdName("standout"),
		pricing.AdName("premium"),
	}
	co := New(pricingRules)
	for _, item := range ads {
		co.Add(item)
	}

	total := co.Total()
	expectedTotal := 269.99 + 322.99 + 394.99
	if total != expectedTotal {
		t.Errorf("found total: %.2f but expected: %.2f", total, expectedTotal)
	}
}
