package main

import (
	"fmt"
	"unicode"
	"bufio"
	"os"
)

func numOfWords(s string) int {
	var runearr []rune = []rune(s)
	var cnt int = 1
	for x := 0; x < len(runearr); x++ {
		if unicode.IsSpace(runearr[x]) {
			cnt++
		}
	}
	for x := 0; x < len(runearr); x++ {
		if unicode.IsUpper(runearr[x]) && x != 0 {
			cnt--
		}
	}
	return cnt
}
func main() {
	var x int = 0
	var word string = ""
	scanner := bufio.NewScanner(os.Stdin)
	var b bool = scanner.Scan()
	for b {
		word += scanner.Text()
		x++
		if x == 1 {
			b = false
		}
	}
	fmt.Println(numOfWords(word))
}
