package main

import (
	"fmt"
)

func main() {
	array := make([]string, 2)
	array = append(array, "1")
	array =




	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)

	}
	balance := calculateBalance(transactions)
	fmt.Printf("Сумма баланса: %.2f", balance)

}

func scanTransaction() float64 {
	var transactions float32
	fmt.Println("Укажите траты(n exit):")
	fmt.Scan(&transactions)
	return float64(transactions)

}

func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}
