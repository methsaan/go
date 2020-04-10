package main

import (
	"fmt"
	"strconv"
	//"math"
	"os"
)

func properFactors(num int) []int {
	var factorList [2000000]int
	var cnt int = 0
	for x := 1.0; int(x) < num; x++ {
		if float64(float64(num)/float64(x)) == float64(int(num)/int(x)) {
			factorList[cnt] = int(x)
			cnt++
		}
	}
	return factorList[0:cnt]
}

func isPerfectNum(num int) bool {
	var propFactors []int = properFactors(num)
	var sum int = 0
	for x := 0; x < len(propFactors); x++ {
		sum += propFactors[x]
	}
	return sum == num
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err == nil {
		fmt.Println("Perfect number:", isPerfectNum(num))
	}
}
