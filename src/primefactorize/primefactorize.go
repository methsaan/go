package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"math/rand"
)

func testEq(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func factors(number int) []int {
	factorList := make([]int, 0)
	for x := 2; x < number; x++ {
		if number%x == 0 {
			factorList = append(factorList, x)
		}
	}
	return factorList
}

func intRepeat(i []int, count int) []int {
	if count == 0 {
		return []int{}
	}
	if count < 0 {
		panic("bytes: negative Repeat count")
	} else if len(i)*count/count != len(i) {
		panic("bytes: Repeat count causes overflow")
	}

	ni := make([]int, len(i)*count)
	bp := copy(ni, i)
	for bp < len(ni) {
		copy(ni[bp:], ni[:bp])
		bp *= 2
	}
	return ni
}

func formatString(num int) string {
	var i string
	if num == -1 {
		i = " "
	} else {
		i = strconv.Itoa(num)
	}
	var remaining int = 5
	var middle int = 2-(len(i)/2)
	var s string = strings.Repeat(" ", middle)
	s = s + i
	remaining -= len(s)
	s = s + strings.Repeat(" ", remaining)
	return s
}

func main() {
	// store prime factors
	primeFactors := make([]int, 50)
	var primeFactorCnt int = 0
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
	var factorTreeRows int = 0
	for x := 1; x < 20; x++ {
		if testEq(factorTree[x-1], append([]int{0}, intRepeat([]int{-1}, len(factorTree[x-1])-1)...)) {
			break
		} else {
			factorTreeRows += 1
		}
		for y := 1; y < pow2[x]; y++ {
			if !factoredSwitch {
				prevRowFactors := factors(factorTree[x-1][(y+1)/2])
				if len(prevRowFactors) == 0 {
					factorPair[0] = -1
					factorPair[1] = -1
					if factorTree[x-1][(y+1)/2] > 1 {
						primeFactors[primeFactorCnt] = factorTree[x-1][(y+1)/2]
						primeFactorCnt++
					}
				} else {
					factorPair[0] = prevRowFactors[rand.Intn(len(prevRowFactors))]
					factorPair[1] = factorTree[x-1][(y+1)/2]/factorPair[0]
				}
			}
			if y%2 != 0 {
				factorTree[x][y] = factorPair[1]
			} else {
				factorTree[x][y] = factorPair[0]
			}
			factoredSwitch = !factoredSwitch
		}
	}
	for x := 0; x < factorTreeRows; x++ {
		fmt.Println(factorTree[x])
	}
	fmt.Println(primeFactors)
	fmt.Println(formatString(-8))
	fmt.Println(formatString(8))
	fmt.Println(formatString(23334))
	fmt.Println(formatString(68))
	fmt.Println(formatString(5))
	fmt.Println(formatString(652))
	fmt.Println(formatString(-1))
}
