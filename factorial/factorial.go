package main

import (
	"fmt"
)

func factorial(num int) int {
	f := num
	for x := num-1; x > 0; x-- {
		f *= x
	}
	return f
}

func sigma(num int, k int) int {
	f := num
	for x := num-1; x > k-1; x-- {
		f += x
	}
	return f
}

func main() {
	var number int
	var start int
	var f int
	var f2 int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &number)
	fmt.Print("Enter \"k\" for sigma: ")
	fmt.Scanf("%d", &start)
	f = factorial(number)
	f2 = sigma(number, start)
	fmt.Println()
	fmt.Printf("%d! = %d\n\n", number, f)
	fmt.Printf("%d\n", number);
	fmt.Printf(" _\n")
	fmt.Printf("  / \tk: %d\n", f2)
	fmt.Printf(" Σ\\\n")
	fmt.Printf("k=%d\n", start)
}
