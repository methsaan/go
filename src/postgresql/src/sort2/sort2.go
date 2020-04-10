package main

import (
	"fmt"
)

func insertionsort(list [100]int, start int, end int) []int {
	var temp int
	for x := start+1; x < end; x++ {
		for y := start; y < end; y++ {
			fmt.Println(list[x], list[y])
			temp = list[x]
			list[x] = list[y]
			list[y] = temp
			fmt.Println(list[start:end])
		}
	}
	return list[start:end]
}

func main() {
	var list [100]int = [100]int{6, 5, 4, 3, 2, 1}
	fmt.Println(list[0:6])
	fmt.Println(insertionsort(list, 0, 6))
}
