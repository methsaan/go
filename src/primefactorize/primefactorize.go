package main

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
)

func factors(number int) []int {
	factorList := make([]int, 0)
	for x := 2; x < number; x++ {
		if number%x == 0 {
			factorList = append(factorList, x)
		}
	}
	return factorList
}

func main() {
	// powers of 2 +1
	pow2 := [20]int{2, 3, 5, 9, 17, 33, 65, 129, 257, 513, 1025, 2049, 4097, 8193, 16385, 32769, 65537, 131073, 262525, 524289}
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
	var factoredSwitch bool = true
	factorPair := make([]int, 2)
	for x := 1; x < 5; x++ {
		for y := 1; y < pow2[x]; y++ {
			if factoredSwitch { // runs every 2 iterations, find factors
				factorPair[0] = // -----------------------------------------------------------------------------------------------------------------------------------
				factorPair[1] = // -----------------------------------------------------------------------------------------------------------------------------------
			}
			if y%2 == 0 { // second factor
				factorTree[x][y] = factorTree[x-1][y/2]
			} else { // first factor
				factorTree[x][y] = factorTree[x-1][(y+1)/2]
			}
			factoredSwitch = !factoredSwitch
		}
	}
	fmt.Println(factorTree[0])
	fmt.Println(factorTree[1])
	fmt.Println(factorTree[2])
	fmt.Println(factorTree[3])
	fmt.Println(factorTree[4])
}
