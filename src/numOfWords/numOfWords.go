package main

import (
	"fmt"
	"unicode"
)

func numOfWords(s string) int {
	var runearr []rune = []rune(s)
	var cnt int = 1
	for x := 0; x < len(runearr); x++ {
		if unicode.IsSpace(runearr[x]) {
			cnt++
		}
	}
	return cnt
}
func main() {
	var word string
	var ho int
	fmt.Print("Enter a word: ")
	word, ho = fmt.Scanln(&word)
	//fmt.Println("Number of words: ", numOfWords(word))
	fmt.Println(ho, word)
}
