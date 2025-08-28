package main

import (
	"fmt"
)

func main() {
	var taxRate float64

	revenue := getUserInput("Enter the total revenue:")

	expenses := getUserInput("Enter the total expenses:")

	taxRate = getUserInput("Enter the tax rate (in %):")

	profitBeforeTax, taxAmount, profitAfterTax, ratio := CalculateProfit(revenue, expenses, taxRate)

	fmt.Printf("Profit before tax: %.2f\n", profitBeforeTax)

	fmt.Printf("Tax amount: %.2f\n", taxAmount)

	fmt.Printf("Profit after tax: %.2f\n", profitAfterTax)

	fmt.Printf("Profit before tax to profit after tax ratio: %.2f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func CalculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64, float64) {
	profitBeforeTax := revenue - expenses
	taxAmount := (taxRate / 100) * profitBeforeTax
	profitAfterTax := profitBeforeTax - taxAmount
	ratio := profitBeforeTax / profitAfterTax
	return profitBeforeTax, taxAmount, profitAfterTax, ratio
}
