package pricing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

// LoadPrices gather all known ad prices, including customer deals
func loadPricesFromJSON() (Rules, error) {
	var priceRules Rules

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	dir = filepath.Dir(dir) //; get parent dir
	data, err := ioutil.ReadFile(dir + "/pricing_rules.json")
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
