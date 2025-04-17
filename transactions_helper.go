package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func calculateBalance(transactions []float64) float64 {
	var balance float64

	for _, value := range transactions {
		balance += value
	}

	return balance
}

func getTransactions(inputReader *bufio.Reader, transactions *[]float64) {
	fmt.Print("Введите транзакции(через пробел): ")

	inputLine, _ := inputReader.ReadString('\n')

	inputLine = strings.TrimSpace(inputLine)
	transactionStrings := strings.Fields(inputLine)

	for _, transactionStr := range transactionStrings {
		transaction, err := strconv.ParseFloat(transactionStr, 64)

		if err != nil {
			fmt.Println("Ошибка при преобразовании транзакции в число:", err)
			continue
		}

		*transactions = append(*transactions, transaction)
	}
}
