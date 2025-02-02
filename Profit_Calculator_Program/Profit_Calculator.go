package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f $\n", ebt)
	fmt.Printf("PROFIT: %.2f $\n", profit)
	fmt.Printf("RATIO: %.2f $\n", ratio)
	storeResults(ebt, profit, ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	//	ebt = earning before tax
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\tProfit: %.2f\tRatio: %.2f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
