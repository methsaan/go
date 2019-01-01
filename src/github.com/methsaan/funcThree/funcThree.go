//check if number is a prime
package main

import (
	"fmt"
)

func isPrime(x int) bool {

	var numberOfFactors,i int = 0,1

	for i <= x {
		if float64(x)/float64(i) == float64(x/i) {
			numberOfFactors++
		}
		i++
	}

	if numberOfFactors == 2 {
		return true
	}else {
		return false
	}

}

func main() {
	var i,cnt int = 0,22
	for i < cnt {
		fmt.Printf("%d is prime: %v\n", i, isPrime(i))
		i++
	}
}
