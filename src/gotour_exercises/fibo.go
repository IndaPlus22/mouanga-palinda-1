package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	num1 := 1
	num2 := 1

	inner_function := func() int {

		result := num1 + num2
		num1 = num2
		num2 = result

		return result

	}

	return inner_function
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
