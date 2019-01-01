package main

import (
	"fmt"
)

func isEven(e int) bool {
	fmt.Printf("Checking if %d is even: ",e)
	return e%2 == 0
}

func isOdd(o int) bool {
	fmt.Printf("Checking if %d is odd: ",o)
	return o%2 == 1
}

func main() {
	fmt.Println(isEven(1))
	fmt.Println(isEven(22))

	fmt.Println(isOdd(10))
	fmt.Println(isOdd(99))
}
