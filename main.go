package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, .07, .1, .15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		var taxIncludedPrices []float64 = make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrices[priceIndex] = getPriceAfterTax(taxRate, price)
		}
		result[taxRate] = taxIncludedPrices
	}

	fmt.Println(result)
}

func getPriceAfterTax(taxRate float64, initialPrice float64) float64 {
	return Round(initialPrice*(1+taxRate), 2)
}

//13.456789, 2

func Round(valToRound float64, precision int) float64 {
	return math.Round(valToRound*math.Pow10(precision)) / math.Pow10(precision)
}
