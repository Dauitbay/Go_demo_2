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
	fmt.Println(transactions)
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Enter the transaction (n for exit): ")
	fmt.Scan(&transaction)
	return transaction
}
