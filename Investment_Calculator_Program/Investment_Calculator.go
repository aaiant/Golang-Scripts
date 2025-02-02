package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Number of years: ")
	fmt.Scan(&years)

	futureValue, realFutureValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value: %.2f.", futureValue)
	formattedFRV := fmt.Sprintf("Future value (adjusted for inflation): %.2f.", realFutureValue)
	fmt.Printf("%s\n%s\n", formattedFV, formattedFRV)
}

//	Functions

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (futureValue, realFutureValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	realFutureValue = futureValue / math.Pow(1+inflationRate/100, years)
	return
}
