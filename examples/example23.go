package main

import (
	"fmt"
)

func print_int(x int) {
	fmt.Println(x)
}

func add_int(x int, y int) int {
	return x + y
}

func sub_int(x int, y int) int {
	return x - y
}

func mul_int(x int, y int) int {
	return x * y
}

func div_int(x int, y int) int {
	return x / y
}

func main() {
	print_int(add_int(2,3))
	print_int(sub_int(4,55))
	print_int(mul_int(4,-5))
	print_int(div_int(-120,-4))
}
