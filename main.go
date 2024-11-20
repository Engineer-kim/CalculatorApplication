package main

import (
	"calculator.com/price-calculator/fileManager"
	"calculator.com/price-calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	for idx, taxRate := range taxRates {
		doneChans[idx] = make(chan bool)
		fm := fileManager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdManager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[idx])
	}
	//모든 goroutine의 작업이 완료될 때까지 메인 goroutine이 대기하는 것
	for _, doneChan := range doneChans {
		<-doneChan
	}
}
