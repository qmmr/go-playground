package main

import (
	"fmt"
	"sync"

	"qmmr.xyz/go/crypto/api"
)

func main() {
	var wg sync.WaitGroup

	currencies := []string{"BTC", "BCH", "ETH"}

	// Loop through all the currencies, create a goroutine for each (similar to async in JS)
	// Use wg sync.WaitGroup to track each goroutine and mark it as done when sync GetRate finishes
	for _, currency := range currencies {
		wg.Add(1)
		// lambdas can be goroutines and to encapsulate currency, they need to be passed as args
		go func(currency string) {
			GetRate(currency)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

func GetRate(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("The rate for \"%v\" is %.2f\n", rate.Currency, rate.Price)
	}
}
