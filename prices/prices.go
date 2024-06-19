package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	Prices            []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) Process() {
	job.TaxIncludedPrices = make(map[string]float64)
	for _, price := range job.Prices {
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = getPriceAfterTax(job.TaxRate, price)
	}

	fmt.Println(job.TaxIncludedPrices)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		Prices:  []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}

func getPriceAfterTax(taxRate float64, initialPrice float64) float64 {
	return initialPrice * (1 + taxRate)
}
