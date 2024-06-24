package main

import (
	"fmt"
	"math"

	"example.com/priceCalculator/filemanager"
	"example.com/priceCalculator/prices"
)

func main() {
	taxRates := []float64{0, .07, .1, .15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("prices-output-%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}

}

func Round(valToRound float64, precision int) float64 {
	return math.Round(valToRound*math.Pow10(precision)) / math.Pow10(precision)
}
