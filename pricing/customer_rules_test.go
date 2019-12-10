package pricing

import "testing"

func getCustomerPriceRules(t *testing.T, customer Customer) CustomerRules {
	rules, err := fetchCustomerRules(customer)
	if err != nil {
		t.Errorf("error when fetching customer ad price rules: %v", err)
	}
	return rules
}

func getPriceForItems(t *testing.T, customer Customer, ads []AdName) float64 {
	return getCustomerPriceRules(t, customer).PriceForItems(ads)
}

func TestDefaultCustomer(t *testing.T) {
	customer := Customer("default")
	customerRules := getCustomerPriceRules(t, customer)
	if len(customerRules.Deals.DiscountPrices) > 0 {
		t.Errorf("expected no discounts for customer '%s'", customer)
	}
	if len(customerRules.Deals.GroupPrices) > 0 {
		t.Errorf("expected no group prices for customer '%s'", customer)
	}
}

func TestStandardAdPricesForCustomer(t *testing.T) {
	customers := []Customer{"SecondBite", "Axil Coffee Roasters", "MYER", "default"}
	for _, customer := range customers {
		customerRules := getCustomerPriceRules(t, customer)
		for ad := range expectedPrices.AdPrices {
			stdPrice := customerRules.AdPrices[ad]
			expectedPrice := expectedPrices.AdPrices[ad]
			if stdPrice != expectedPrices.AdPrices[ad] {
				t.Errorf("customer '%s' ad price '%s' found: %.2f but expected: %.2f", customer, ad, stdPrice, expectedPrice)
			}
		}
	}
}

func TestStandardAdPricesMultiple(t *testing.T) {
	customer := Customer("unknown customer")
	for ad := range expectedPrices.AdPrices {
		price := getPriceForItems(t, customer, []AdName{ad, ad, ad})
		expectedClassicPrice := 3 * expectedPrices.AdPrices[ad]
		if price != expectedClassicPrice {
			t.Errorf("customer '%s' ad '%s' price found: %.2f but expected: %.2f", customer, ad, price, expectedClassicPrice)
		}
	}
}

func TestGroupPrices(t *testing.T) {
	// SecondBite has a 3-for-2 offer on "classic"
	customer := Customer("SecondBite")
	ad := AdName("classic")
	rules := getCustomerPriceRules(t, customer)
	adPrice := rules.AdPrices[ad]

	price := getPriceForItems(t, customer, []AdName{ad, ad, ad})
	expectedPrice := 2 * adPrice
	if price != expectedPrice {
		t.Errorf("customer '%s' ad '%s' group price found: %.2f but expected: %.2f", customer, ad, price, expectedPrice)
	}

	// test multiples of the group count
	price = getPriceForItems(t, customer, []AdName{ad, ad, ad, ad, ad, ad, ad, ad, ad})
	expectedPrice = 3 * 2 * adPrice
	if price != expectedPrice {
		t.Errorf("customer '%s' ad '%s' group price found: %.2f but expected: %.2f", customer, ad, price, expectedPrice)
	}

	// MYER has a 5-for-4 offer on "classic"
	customer = Customer("MYER")
	ad = AdName("classic")
	rules = getCustomerPriceRules(t, customer)
	adPrice = rules.AdPrices[ad]

	price = getPriceForItems(t, customer, []AdName{ad, ad, ad, ad, ad})
	expectedPrice = 4 * adPrice
	if price != expectedPrice {
		t.Errorf("customer '%s' ad '%s' group price found: %.2f but expected: %.2f", customer, ad, price, expectedPrice)
	}

	// test slightly more than the group count
	// -- the final two ads should be given the standard price
	price = getPriceForItems(t, customer, []AdName{ad, ad, ad, ad, ad, ad, ad})
	expectedPrice = 4*adPrice + 2*adPrice
	if price != expectedPrice {
		t.Errorf("customer '%s' ad '%s' group price found: %.2f but expected: %.2f", customer, ad, price, expectedPrice)
	}
}

func TestDiscountAdPrices(t *testing.T) {
	customer := Customer("Axil Coffee Roasters")
	ad := AdName("standout")

	price := getPriceForItems(t, customer, []AdName{ad})
	expectedPrice := expectedPrices.Deals[customer].DiscountPrices[ad]
	if price != expectedPrice {
		t.Errorf("customer '%s' ad '%s' discount price found: %.2f but expected: %.2f", customer, ad, price, expectedPrice)
	}

	customer = Customer("MYER")
	ad = AdName("premium")

	price = getPriceForItems(t, customer, []AdName{ad})
	expectedPrice = expectedPrices.Deals[customer].DiscountPrices[ad]
	if price != expectedPrice {
		t.Errorf("customer '%s' ad '%s' discount price found: %.2f but expected: %.2f", customer, ad, price, expectedPrice)
	}
}

// we assume group and discount pricing deals do not overlap
// e.g. if customer SecondBite has a 3-for-2 deal on Classic Ads
// then they will *not* also have a discount on Classic Ads
//
// vice versa, if customer MYER has a 5-for-4 deal on Premium Ads
// then they will *not* also have a discount on Premium Ads

func TestMultipleAdPricings(t *testing.T) {
	customer := Customer("MYER")
	rules := getCustomerPriceRules(t, customer)

	price := getPriceForItems(t, customer, []AdName{"classic", "classic", "classic", "classic", "classic", "premium", "standout"})

	// 5-for-4 classic + 1 premium discount + 1 standout standard
	expectedPrice := 4.0*rules.AdPrices[AdName("classic")] + rules.Deals.DiscountPrices[AdName("premium")] + rules.AdPrices[AdName("standout")]
	if price != expectedPrice {
		t.Errorf("customer '%s' total price found: %.2f but expected: %.2f", customer, price, expectedPrice)
	}
}

func TestExampleScenarios(t *testing.T) {
	customer := Customer("default")
	price := getPriceForItems(t, customer, []AdName{"classic", "standout", "premium"})
	expectedPrice := 987.97
	if price != expectedPrice {
		t.Errorf("customer '%s' total price found: %.2f but expected: %.2f", customer, price, expectedPrice)
	}

	customer = Customer("SecondBite")
	price = getPriceForItems(t, customer, []AdName{"classic", "classic", "classic", "premium"})
	expectedPrice = 934.97
	if price != expectedPrice {
		t.Errorf("customer '%s' total price found: %.2f but expected: %.2f", customer, price, expectedPrice)
	}

	customer = Customer("Axil Coffee Roasters")
	price = getPriceForItems(t, customer, []AdName{"standout", "standout", "standout", "premium"})
	expectedPrice = 1294.96
	if price != expectedPrice {
		t.Errorf("customer '%s' total price found: %.2f but expected: %.2f", customer, price, expectedPrice)
	}
}
