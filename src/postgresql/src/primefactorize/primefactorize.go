package main

import (
	"fmt"
	"rand"
)

func main() {
	var number int
	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &number)
	var fullFactors [100]int
	var numOfFactors int
	for x := 1.0; int(x) <= number; x += 1.0 {
		if float64(number)/float64(x) == float64(number/int(x)) {
			primeFactors[numOfFactors] = int(x)
			numOfFactors += 1
		}
	}
	newFactors = fullFactors[0:numOfFactors]
	for true {
	}
}
