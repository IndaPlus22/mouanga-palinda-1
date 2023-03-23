package main

import "fmt"

func Sqrt(x float64) float64 {
	// intial variables

	iterations := 0

	result := (0.0 + x) / 2

	for iterations < 10 {

		result -= (result*result - x) / (2 * result)
		iterations += 1

	}

	return result
}

func main() {
	fmt.Println("Sqrt(500) ~=", Sqrt(500))
}
