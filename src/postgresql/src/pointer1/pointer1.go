package main

import "fmt"

func main() {
	var x int = 10
	var y int
	var px *int = &x
	y = *px
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("px:", px)
	fmt.Println("x == y:", x==y)
}
