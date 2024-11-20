package prices

import (
	"calculator.com/price-calculator/conversion"
	"calculator.com/price-calculator/iomnager"
	"fmt"
)

// 단순히 구조 정의 즉. 타입 정의
type TaxIncludedPriceJob struct {
	IOManager         iomnager.IOManager `json:"-"`
	TaxRate           float64            `json:"tax_rate"`
	InputPrices       []float64          `json:"input_prices"`
	TaxIncludedPrices map[string]string  `json:"tax_included_prices"`
}

// 솟값을 전달안하면  실제 구조체로 전달된 필드의값을 변경할수 없기 떄문에  TaxIncludedPriceJob 의 값이 아닌 주솟값전달(얕은 복사)
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	//Sprintf : 문자열을 포맷팅
	for _, price := range job.InputPrices {
		taxIncludePrice := price * (1 + job.TaxRate)
		//result[키] = 벨류
		result[fmt.Sprintf("%.2f 일때", price)] = fmt.Sprintf("%.2f", taxIncludePrice)
	}

	//fmt.Println(result)
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
}

// 컨스트럭션 함수  (메모리 할당)
func NewTaxIncludedPriceJob(iom iomnager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		//TaxIncludedPrices: make(map[string]float64),
	}
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	//for lineIndex, line := range lines {
	//	floatPrice, err := strconv.ParseFloat(line, 64)
	//
	if err != nil {
		fmt.Println("Converting price to Float falied")
		fmt.Println(err)
		return err
	}
	//
	//	prices[lineIndex] = floatPrice
	//}

	job.InputPrices = prices
	return nil
}
