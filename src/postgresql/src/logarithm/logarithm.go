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

func declog(b float64, n float64) float64 {
	return math.Log(n)/math.Log(b)
}

func main() {
	var base float64
	var num float64
	fmt.Println("Enter the base: ")
	fmt.Scanf("%f", &base)
	fmt.Println("Enter a number: ")
	fmt.Scanf("%f", &num)
	fmt.Println("Integer logarithm: ", logarithm(base, num))
	fmt.Println("Decimal logarithm: ", declog(base, num))
}
