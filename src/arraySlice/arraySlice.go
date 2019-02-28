package main

import (
	"fmt"
)

func getstr(str *[]rune) {
	fmt.Scanf("%s", string(*str))
}
func main() {
	var str[8]rune
	var stringSlice = str[:]
	fmt.Print("Enter a string: ")
	getstr(&stringSlice)
	fmt.Println(str)
}
