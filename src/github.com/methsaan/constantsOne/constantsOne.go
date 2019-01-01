package main

import (
	"fmt"
)

// this function never gets the addresses of a and b
func changeValueCons(x *int ,y *int) *int {
	*x = 30
	*y = 40
	var z int = *x+*y
	return(&z)
}

func changeValueVar(m *int, n  *int) *int {
	*m = 120
	*n = 220
	var o int = *m+*n
	return(&o)
}

func main() {
	const a int = 10
	const b int = 20
	fmt.Println(a+b)
	// cannot take address of constant - un-comment to test/see error
	// fmt.Println(changeValueCons(&a,&b))
	//ptra := &a
	//ptrb := &b

	var c int = 100
	var d int = 200
	fmt.Println(c+d)
	fmt.Println(*changeValueVar(&c,&d))
}
