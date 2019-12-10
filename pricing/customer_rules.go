package pricing

// FetchPricingRules all pricing rules and deals associated with the given customer
func FetchPricingRules(customer Customer) (CustomerAdPricer, error) {
	return fetchCustomerRules(customer)
}

// PriceForItems calculate the price of the given ads for the given customer
func (cr CustomerRules) PriceForItems(ads []AdName) float64 {
	adCounts := cr.countAllAdNames(ads)
	adPrices := make(map[AdName]float64, len(adCounts))

	for ad := range adCounts {
		if cr.hasGroupPriceForAd(ad) {
			adPrices[ad] = cr.groupPriceForAd(ad, adCounts[ad])
		} else if cr.hasDiscountPriceForAd(ad) {
			adPrices[ad] = cr.discountPriceForAd(ad) * float64(adCounts[ad])
		} else {
			adPrices[ad] = cr.standardPriceForAd(ad) * float64(adCounts[ad])
		}
	}

	var totalPrice float64
	for _, price := range adPrices {
		totalPrice += price
	}
	return totalPrice
}

// fetchRules get the current state of all pricing rules
// Ordinarily this might be sourced from a service,
// but for now let's just load this from a JSON file
func fetchRules() (Rules, error) {
	return loadPricesFromJSON()
}

// fetchCustomerRules get all pricing rules for a given Customer
func fetchCustomerRules(customer Customer) (CustomerRules, error) {
	var customerRules CustomerRules
	allRules, err := fetchRules()
	if err != nil {
		return customerRules, err
	}

	_, ok := allRules.Deals[customer]
	if ok {
		// rules and deals for known customer
		return CustomerRules{
			CustomerName: customer,
			AdPrices:     allRules.AdPrices,
			Deals:        allRules.Deals[customer],
		}, nil
	}
	// rules for 'generic' customer, i.e. std rules, no deals
	return CustomerRules{
		CustomerName: customer,
		AdPrices:     allRules.AdPrices,
		Deals:        CustomerPricingRules{},
	}, nil
}

// knownAds return an array of all known AdNames
func (cr CustomerRules) knownAdNames() []AdName {
	keys := make([]AdName, 0, len(cr.AdPrices))
	for k := range cr.AdPrices {
		keys = append(keys, k)
	}
	return keys
}

func (cr CustomerRules) hasGroupPriceForAd(ad AdName) bool {
	_, ok := cr.Deals.GroupPrices[ad]
	return ok
}

func (cr CustomerRules) groupPriceForAd(ad AdName, adCount int) float64 {
	adGroupPrice := cr.Deals.GroupPrices[ad]

	// in a 3-for-2 deal, groupCount == 3, multiplier == 2
	groupCount := adGroupPrice.GroupCount
	multiplier := adGroupPrice.Multiplier

	groups := float64(adCount / groupCount)
	remainder := float64(adCount % groupCount)
	stdPrice := cr.standardPriceForAd(ad)

	return (groups * float64(multiplier) * stdPrice) + (remainder * stdPrice)
}

func (cr CustomerRules) hasDiscountPriceForAd(ad AdName) bool {
	_, ok := cr.Deals.DiscountPrices[ad]
	return ok
}

func (cr CustomerRules) discountPriceForAd(ad AdName) float64 {
	return cr.Deals.DiscountPrices[ad]
}

func (cr CustomerRules) standardPriceForAd(ad AdName) float64 {
	return cr.AdPrices[ad]
}

// countAds return the number of times ad occurs in ads
func countAds(ad AdName, ads []AdName) int {
	adCount := 0
	for _, thisAdName := range ads {
		if thisAdName == ad {
			adCount++
		}
	}
	return adCount
}

// countAllAds the count of all AdNames in ads
func (cr CustomerRules) countAllAdNames(ads []AdName) map[AdName]int {
	adNames := cr.knownAdNames()
	adCounts := make(map[AdName]int, len(adNames))

	for _, ad := range adNames {
		adCounts[ad] = countAds(ad, ads)
	}
	return adCounts
}
