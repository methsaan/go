//check if number is a prime
package main

import (
	"fmt"
)

func isPrime(x int) bool {
	var numberOfFactors,i int = 0,1
	for i < x {
		if float64(x)/float64(i) == float64(x/i) {
			numberOfFactors++
		}
		i++
	}
	numberOfFactors++
	if numberOfFactors == 2 {
		return true
	}else {
		return false
	}
}

func main() {
	var i int = 0

	numbers := [5]int{2,16,5,35,2797}
	for i < 5 {
		fmt.Printf("%d is prime: %v\n", numbers[i], isPrime(numbers[i]))
		i++
	}
}
