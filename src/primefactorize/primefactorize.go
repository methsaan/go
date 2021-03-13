package main

import (
	"fmt"
	"os"
	"bytes"
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
	slice1 := make([]int, 3)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	res1 := bytes.Repeat(slice1, 3)
	fmt.Printf("Result: %s\n", res1)
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
	var factoredSwitch bool = false
	factorPair := make([]int, 2)
	for x := 1; x < 6; x++ {
		// see if previous row is equal to [0] + ([-1] * length of previous row)
		fmt.Println(factorTree[x-1])
		//if factorTree[x-1] {
		//	
		//}
		for y := 1; y < pow2[x]; y++ {
			if !factoredSwitch { // runs every 2 iterations, find factors
				prevRowFactors := factors(factorTree[x-1][(y+1)/2])
				if len(prevRowFactors) == 0 {
					factorPair[0] = -1
					factorPair[1] = -1
				} else {
					factorPair[0] = prevRowFactors[rand.Intn(len(prevRowFactors))]
					factorPair[1] = factorTree[x-1][(y+1)/2]/factorPair[0]
				}
			}
			if y%2 != 0 { // first factor
				factorTree[x][y] = factorPair[1]
			} else { // second factor
				factorTree[x][y] = factorPair[0]
			}
			factoredSwitch = !factoredSwitch
		}
	}
	fmt.Println(factorTree[0])
	fmt.Println(factorTree[1])
	fmt.Println(factorTree[2])
	fmt.Println(factorTree[3])
	fmt.Println(factorTree[4])
	fmt.Println(factorTree[5])
}
