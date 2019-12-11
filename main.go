package main

import (
	"fmt"

	"github.com/dhodges/scc/checkout"
	"github.com/dhodges/scc/pricing"
)

func printCheckout(customer pricing.Customer, co checkout.Checkout) {
	fmt.Printf("Customer: %s\n", customer)
	fmt.Print("items:   ")
	for _, item := range co.Items() {
		fmt.Printf("`%s`,", item)
	}
	fmt.Println()
	fmt.Printf("Total:   $%.2f\n", co.Total())
	fmt.Println()
}

func checkoutAndPrint(customer pricing.Customer, ads []pricing.AdName) {
	pricingRules, err := pricing.FetchPricingRules(customer)
	if err != nil {
		fmt.Printf("error when fetching pricing rules: %v", err)
	} else {
		co := checkout.New(pricingRules)
		for _, ad := range ads {
			co.Add(ad)
		}
		printCheckout(customer, co)
	}
}

func main() {

	fmt.Println()
	fmt.Println("Example scenarios")
	fmt.Println()

	checkoutAndPrint(pricing.Customer("default"), []pricing.AdName{
		pricing.AdName("classic"),
		pricing.AdName("standout"),
		pricing.AdName("premium"),
	})

	checkoutAndPrint(pricing.Customer("SecondBite"), []pricing.AdName{
		pricing.AdName("classic"),
		pricing.AdName("classic"),
		pricing.AdName("classic"),
		pricing.AdName("premium"),
	})

	checkoutAndPrint(pricing.Customer("Axil Coffee Roasters"), []pricing.AdName{
		pricing.AdName("standout"),
		pricing.AdName("standout"),
		pricing.AdName("standout"),
		pricing.AdName("premium"),
	})
}
