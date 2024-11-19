package prices

import (
	"bufio"
	"fmt"
	"os"
)

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

func (job TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could not open prices.txt")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading The FIle content Falied")
		fmt.Println(err)
		file.Close()
		return
	}
}
