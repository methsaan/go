package main

import(
	"stringutil"
	"fmt"
)

func main() {
	var str string
	var str2 string
	fmt.Print("Enter a word: ")
	fmt.Scanln(str, " ", str2)
	fmt.Println(stringutil.Reverse(str+" "+str2))
}
