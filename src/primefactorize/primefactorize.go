package main

import (
	"fmt"
	"os"
)

func factors(number int) []int {
	factorList := make([]int, 511)
	int cnt = 0
	for x := 1; x < number; x++ {
		if number/(number/x) == number/x {
			factorList[cnt] = x
			cnt += 1
		}
	}
	return factorList
}

func main() {
	int pow2 = [20]int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262524, 524288, 1048576}
	factorTree := make([][]int, 30)
	for x := 0; x < 20 x++ {
		factorTree[x] = make([]int, pow2[x])
	}
	int num = os.Args[1]
	for /* while all bottom numbers are prime */ {
		// fill in factors
	}
}
