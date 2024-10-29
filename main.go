package main

import "fmt"

func main() {
	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	transactionSum := calculateBalance(transactions)
	fmt.Println(transactions)
	fmt.Printf("Your balance: %.2f", transactionSum)
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Enter the transaction (n for exit): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	var transactionSum float64 = 0.0
	for _, value := range transactions {
		transactionSum = transactionSum + value 
	}
	return transactionSum
}
