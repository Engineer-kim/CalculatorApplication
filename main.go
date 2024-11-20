package main

import (
	"calculator.com/price-calculator/fileManager"
	"calculator.com/price-calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	for _, taxRate := range taxRates {
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdManager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
