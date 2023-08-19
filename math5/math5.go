package main

import (
	"fmt"
)

func mem_add(a int) *int {
	xa := &a
	return xa
}

func mem_val(b *int) int {
	return *b
}

func mem_add_rand(a int) int {
	ap := &a
	return *(ap+10)
}

func main() {
	var y int = 33
	var z int = 55
	fmt.Println("Mem address",mem_add(y))
	fmt.Println("Value",mem_val(&z))
	fmt.Println("Mem address rand",mem_add_rand(y))
}
