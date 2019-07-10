package main

import(
	"fmt"
	"math"
)

func isIn(float64 arr[], float64 num){
}
func logarithm(b float64, n float64) int {
	var p int = 0
	for x := 0; !(math.Abs(n-1.000000) < 0.3) && x < 32767; x++{
		n /= b
		p++
	}
	return p
}

func declog(b float64, n float64) float64 {
	var p float64 = 0.0000000000000
	var powers []uint = uint[1000]
	for x := 0; x < 1000; x++ {
		powers[x] = math.Pow(b, x)
	}
	for x := 0; !(math.Abs(math.Pow(b, p)-n) < 0.00001); x++ {
		p += 0.00001
		if (isIn(power, math.Pow(b, p))) {
			fmt.Println("Loss of precision")
		}
		fmt.Printf("%.5f\n", p)
	}
	return p+0.00000001
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
