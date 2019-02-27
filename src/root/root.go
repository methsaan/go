package main

import (
	"fmt"
	"math"
)

func main() {
	var base float64
	var root float64
	fmt.Print("Enter a number: ")
	fmt.Scanf("%f", &base)
	fmt.Print("Enter the root: ")
	fmt.Scanf("%f", &root)
	fmt.Printf("%f root %f = %.20f\n", base, root, math.Pow(base, 1/root))
}
