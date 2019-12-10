package pricing

import (
	"testing"
)

var expectedPrices = Rules{
	AdPrices: map[AdName]float64{
		"classic":  float64(269.99),
		"standout": float64(322.99),
		"premium":  float64(394.99),
	},
	Deals: map[Customer]CustomerPricingRules{
		"SecondBite": CustomerPricingRules{
			DiscountPrices: map[AdName]float64{},
			GroupPrices: map[AdName]AdGroupPrice{
				"classic": AdGroupPrice{
					GroupCount: int(3),
					Multiplier: int(2),
				},
			},
		},
		"Axil Coffee Roasters": CustomerPricingRules{
			DiscountPrices: map[AdName]float64{
				"standout": float64(299.99),
			},
			GroupPrices: map[AdName]AdGroupPrice{},
		},
		"MYER": CustomerPricingRules{
			DiscountPrices: map[AdName]float64{
				"premium": float64(389.99),
			},
			GroupPrices: map[AdName]AdGroupPrice{
				"classic": AdGroupPrice{
					GroupCount: int(5),
					Multiplier: int(4),
				},
			},
		},
	},
}

func TestLoadJSON(t *testing.T) {
	loadedPrices, err := loadPricesFromJSON()
	if err != nil {
		t.Errorf("loading JSON returned error: %v", err)
	}

	for ad := range expectedPrices.AdPrices {
		price := loadedPrices.AdPrices[ad]
		expectedPrice := expectedPrices.AdPrices[ad]
		if price != expectedPrice {
			t.Errorf("AdPrices: ad '%s' is: %f, but expected: %f", ad, price, expectedPrice)
		}
	}

	for customer := range expectedPrices.Deals {
		rules := loadedPrices.Deals[customer]
		expectedRules := expectedPrices.Deals[customer]

		for ad := range expectedRules.DiscountPrices {
			discount := rules.DiscountPrices[ad]
			expectedDiscount := expectedRules.DiscountPrices[ad]
			if discount != expectedDiscount {
				t.Errorf("Deals: customer: '%s' ad: '%s' discount is: %f but expected: %f", customer, ad, discount, expectedDiscount)
			}
		}

		for ad := range expectedRules.GroupPrices {
			groupPrices := rules.GroupPrices[ad]
			expectedGroupPrices := expectedRules.GroupPrices[ad]

			groupCount := groupPrices.GroupCount
			expectedGroupCount := expectedGroupPrices.GroupCount
			if groupCount != expectedGroupCount {
				t.Errorf("Deals: customer: '%s' ad: '%s' group count is: %d but expected: %d", customer, ad, groupCount, expectedGroupCount)
			}

			multiplier := groupPrices.Multiplier
			expectedMultiplier := expectedGroupPrices.Multiplier
			if multiplier != expectedMultiplier {
				t.Errorf("Deals: customer: '%s' ad: '%s' multiplier is: %d but expected: %d", customer, ad, multiplier, expectedMultiplier)
			}
		}
	}
}
