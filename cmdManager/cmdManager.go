package cmdManager

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Enter value")

	var prices []string

	for {
		var price string
		fmt.Print("Enter price: ")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}