package pricing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// LoadPrices gather all known ad prices, including customer deals
func loadPricesFromJSON() (Rules, error) {
	var priceRules Rules

	data, err := ioutil.ReadFile("./pricing_rules.json")
	if err != nil {
		fmt.Print(err)
	} else {
		err = json.Unmarshal(data, &priceRules)
		if err != nil {
			fmt.Print(err)
		}
	}
	return priceRules, err
}
