package main

import (
	"math"

	"example.com/priceCalculator/prices"
)

func main() {
	taxRates := []float64{0, .07, .1, .15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

}

//13.456789, 2

func Round(valToRound float64, precision int) float64 {
	return math.Round(valToRound*math.Pow10(precision)) / math.Pow10(precision)
}
