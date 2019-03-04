package main

import(
	"fmt"
	"math"
)

func logarithm(b float64, n float64) int {
	var p int = 0
	for x := 0; !(math.Abs(n-1.000000) < 0.3) && x < 32767; x++{
		n /= b
		p++
	}
	return p
}

func main() {
	var num float64
	var base float64
	fmt.Println("Enter a number: ")
	fmt.Scanf("%f", &num)
	fmt.Println("Enter the base: ")
	fmt.Scanf("%f", &base)
	fmt.Println(logarithm(base, num))
}
