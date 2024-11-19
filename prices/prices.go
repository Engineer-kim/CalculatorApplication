package prices

import "fmt"

// 단순히 구조 정의 즉. 타입 정의
type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)

	//Sprintf : 문자열을 포맷팅
	for _, price := range job.InputPrices {
		//result[키] = 벨류
		result[fmt.Sprintf("%.2f::", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)
}

// 컨스트럭션 함수  (메모리 할당)
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		//TaxIncludedPrices: make(map[string]float64),
	}
}
