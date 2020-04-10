package main

import (
	"fmt"
	"unicode"
	"strings"
)

func main() {
	var s1 string = "Hello"
	fmt.Println(strings.ToUpper(s1))
	var r1 rune = '4'
	var str = [12]rune{'H', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd', '.'}
	fmt.Println("unicode.IsDigit(r1)", unicode.IsDigit(r1))
	fmt.Println("strings.ToUpper(s1)", strings.ToUpper(s1))
	fmt.Println("unicode.IsUpper(str[0])", unicode.IsUpper(str[1]))
	fmt.Println("unicode.IsPunct(str[11])", unicode.IsPunct(str[11]))
	fmt.Println("unicode.IsSpace(str[12])", unicode.IsSpace(str[5]))
	for x := 0; x < 12; x++ {
		fmt.Printf("%c", unicode.ToUpper(str[x]))
	}
	fmt.Println()
}
