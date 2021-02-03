package main

import (
	"fmt"
	"os"
	"strconv"
)

func factors(number int) []int {
	factorList := make([]int, 511)
	var cnt int = 0
	for x := 1; x < number; x++ {
		if number/(number/x) == number/x {
			factorList[cnt] = x
			cnt += 1
		}
	}
	return factorList
}

func main() {
	pow2 := [20]int{2, 3, 5, 9, 17, 33, 65, 129, 257, 513, 1025, 2049, 4097, 8193, 16385, 32769, 65537, 131073, 262525, 524289}
	// powers of 2 +1
	factorTree := make([][]int, 20)
	for x := 0; x < 20; x++ {
		factorTree[x] = make([]int, pow2[x])
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	for x := 0; x < 20; x++ {
		factorTree[x][0] = 0
	}
	factorTree[0][1] = num
	var cnt int = 0
	for x := 1; x < 5; x++ {
		for y := 1; y < pow2[x]; y++ {
			if y%2 == 0 {
				factorTree[x][y] = factorTree[x-1][y] * 2 - 1
				//index out of range
			} else {
				factorTree[x][y] = factorTree[x-1][y] * 2 + 1
			}
			cnt++
		}
	}
	fmt.Println(factorTree[0])
	fmt.Println(factorTree[1])
	fmt.Println(factorTree[2])
	fmt.Println(factorTree[3])
	fmt.Println(factorTree[4])
}
