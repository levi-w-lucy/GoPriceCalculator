package main

import (
	"fmt"
	"math"

	"example.com/priceCalculator/filemanager"
	"example.com/priceCalculator/prices"
)

func main() {
	taxRates := []float64{0, .07, .1, .15}
	doneChans := make([]chan bool, len(taxRates))

	for taxRateIdx, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("prices-output-%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		doneChans[taxRateIdx] = make(chan bool)
		go priceJob.Process(doneChans[taxRateIdx])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// 	return
		// }
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}

}

func Round(valToRound float64, precision int) float64 {
	return math.Round(valToRound*math.Pow10(precision)) / math.Pow10(precision)
}
