package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	OutputText("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	OutputText("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expected annual return rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future Value: %.2f\n", futureValue)
	realValue := fmt.Sprintf("Future Real Value: %.2f", futureRealValue)

	fmt.Println(realValue)	
}

func OutputText(text string) {
	fmt.Print(text)
}
