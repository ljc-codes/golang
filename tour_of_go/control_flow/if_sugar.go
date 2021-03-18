package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if value := math.Pow(x, n); value < lim {
		return value
	}
	return lim
}

func main() {
	fmt.Println(pow(2, 5, 19))         // prints 19, hits lim
	fmt.Println(pow(2, 2, 1900000000)) // prints 4, ~ hits lim
}
