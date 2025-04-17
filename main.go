package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("__Транзакции__")
	var transactions []float64

	transactionsLen := 0
	balance := 0.0

	userChoise := 1
	inputReader := bufio.NewReader(os.Stdin)

	for {
		if userChoise == 1 {
			getTransactions(inputReader, &transactions)
		} else if userChoise == 2 {
			fmt.Print("Транзакции: ")
			fmt.Println(transactions)
		} else if userChoise == 3 {
			break
		} else {
			fmt.Println("Неверный ввод")
		}
		fmt.Println("1. Добавить транзакции\n2. Посмотреть транзакции\n3. Выход")

		if transactionsLen != len(transactions) {
			balance = calculateBalance(transactions)
			fmt.Printf("Баланс: %.2f\n", balance)
		} else {
			fmt.Printf("Баланс: %.2f\n", balance)
		}

		fmt.Scan(&userChoise)
		inputReader.ReadString('\n')
	}
}
