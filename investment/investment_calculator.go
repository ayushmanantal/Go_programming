package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	OutputText("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	OutputText("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expected annual return rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := CalculateFutureValue(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Future Value: %.2f\n", futureValue)
	realValue := fmt.Sprintf("Future Real Value: %.2f", futureRealValue)

	fmt.Println(realValue)
}

func OutputText(text string) {
	fmt.Print(text)
}

func CalculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
