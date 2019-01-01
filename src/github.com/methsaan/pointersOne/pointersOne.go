package main

import "fmt"

func main() {
	var a int = 30
	var b* int = &a
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b)
}
