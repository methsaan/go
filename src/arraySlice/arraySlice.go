package main

import (
	"fmt"
)

func getstr() []rune {
	var s string
	fmt.Scanf("%s", &s)
	return []rune(s)
}
func printstr(str []rune) {
	for x := 0; x < len(str); x++ {
		fmt.Printf("%c", str[x])
	}
	fmt.Println()
}
func main() {
	fmt.Print("Enter a string: ")
	str := getstr()
	printstr(str)
}
