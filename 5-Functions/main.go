package main

import (
	"fmt"
)

func main() {
	johnPrice := computePrice(145.90, 3)
	paulPrice := computePrice(26.32, 10)

	total := johnPrice + paulPrice

	fmt.Printf("TOTAL : $ %0.2f \n", total)
}

func computePrice(rate float32, night int) float32 {
	return rate * float32(night)
}