package prices

import (
	"bufio"
	"fmt"
	"os"

	"example.com/priceCalculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	Prices            []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadPrices() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could not read file prices.txt")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var prices []string
	for scanner.Scan() {
		prices = append(prices, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading file content failed")
		fmt.Println(err)
		file.Close()
		return
	}

	job.Prices, err = conversion.StringsToFloats(prices)
	if err != nil {
		fmt.Println(err)
		return
	}

	file.Close()
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrices()

	job.TaxIncludedPrices = make(map[string]string)
	for _, price := range job.Prices {
		taxIncludedPrice := getPriceAfterTax(job.TaxRate, price)
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(job.TaxIncludedPrices)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func getPriceAfterTax(taxRate float64, initialPrice float64) float64 {
	return initialPrice * (1 + taxRate)
}
