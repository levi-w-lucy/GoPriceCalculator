package prices

import (
	"fmt"

	"example.com/priceCalculator/conversion"
	"example.com/priceCalculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	Prices            []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadPrices() error {
	prices, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	job.Prices, err = conversion.StringsToFloats(prices)
	if err != nil {
		return err
	}

	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadPrices()
	if err != nil {
		return err
	}

	job.TaxIncludedPrices = make(map[string]string)
	for _, price := range job.Prices {
		taxIncludedPrice := getPriceAfterTax(job.TaxRate, price)
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: fm,
		TaxRate:   taxRate,
	}
}

func getPriceAfterTax(taxRate float64, initialPrice float64) float64 {
	return initialPrice * (1 + taxRate)
}
