package list

import "fmt"

func List() {
	prices := []float64{
		10.99,
		8.99,
	}
	prices[1] = 9.99

	prices = append(prices, 5.99, 6.8, 12.99, 100.10)
	prices = prices[1:]
	fmt.Println(prices)
	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)

}
