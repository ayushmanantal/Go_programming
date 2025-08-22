package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Println("Enter the total revenue:")
	fmt.Scan(&revenue)

	fmt.Println("Enter the total expenses:")
	fmt.Scan(&expenses)

	fmt.Println("Enter the tax rate (as a percentage):")
	fmt.Scan(&taxRate)

	// Calculate profit before tax
	profitBeforeTax := revenue - expenses
	fmt.Printf("Profit before tax: %.2f\n", profitBeforeTax)

	// Calculate tax amount
	taxAmount := (taxRate / 100) * profitBeforeTax
	fmt.Printf("Tax amount: %.2f\n", taxAmount)

	// Calculate profit after tax
	profitAfterTax := profitBeforeTax - taxAmount
	fmt.Printf("Profit after tax: %.2f\n", profitAfterTax)

	// Calculate ratio
	ratio := profitBeforeTax / profitAfterTax
	fmt.Printf("Profit before tax to profit after tax ratio: %.2f\n", ratio)
}
