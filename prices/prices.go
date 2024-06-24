package prices

import (
	"fmt"

	"example.com/priceCalculator/conversion"
	"example.com/priceCalculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	Prices            []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadPrices() {
	prices, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	job.Prices, err = conversion.StringsToFloats(prices)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrices()

	job.TaxIncludedPrices = make(map[string]string)
	for _, price := range job.Prices {
		taxIncludedPrice := getPriceAfterTax(job.TaxRate, price)
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	filemanager.WriteJSON(fmt.Sprintf("prices-output-%.0f.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func getPriceAfterTax(taxRate float64, initialPrice float64) float64 {
	return initialPrice * (1 + taxRate)
}
