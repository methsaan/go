//check if number is a prime
package main

import (
	"fmt"
)

func isPrime(x int) bool {
	var numberOfFactors,i,j int = 0,1,0
	var factors [100]int
	for i < x {
		if float64(x)/float64(i) == float64(x/i) {
			numberOfFactors++
			factors[j] = i
			j++
		}
		i++
	}
	numberOfFactors++
	fmt.Println(factors)
	if numberOfFactors == 2 {
		return true
	}else {
		return false
	}
}

func main() {
	var i,c,cnt int = 0,0,10000
	for c < cnt {
		fmt.Printf("%d is prime: %v\n", c, isPrime(c))
		c++
	}

	numbers := [5]int{2,16,5,35,79}
	for i < 5 {
		fmt.Printf("%d is prime: %v\n", numbers[i], isPrime(numbers[i]))
	i++
	}
}
