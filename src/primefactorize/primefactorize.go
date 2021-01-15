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
	pow2 := [20]int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262524, 524288, 1048576}
	factorTree := make([][]int, 20)
	for x := 0; x < 20; x++ {
		factorTree[x] = make([]int, pow2[x])
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	factorTree[0][0] = num
	var cnt int = 0
	for x := 0; x < 5; x++ {
		for y := 0; y < pow2[x]; y++ {
			factorTree[x][y] = cnt
			cnt++
		}
	}
	fmt.Println(factorTree[0:5])
}
