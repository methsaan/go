package main

import (
	"fmt"
)

func sayString(s string) string {
	return "Hello " + s
}

func sayInt(i int) int {
	return i
}

func sayFloat32(f float32) float32 {
	return f
}

func sayFloat64(ff float64) float64 {
	return ff
}

func sayBool(b bool) bool {
	return b
}

func sayBigInt(ii uint) uint {
	return ii
}

/*
func sayIntArray(ia int) int {
	return ia
}
*/

func main() {
	fmt.Println(sayString("George"))
	fmt.Println(sayInt(12))
        fmt.Println(sayFloat32(86.796536))
	fmt.Println(sayFloat64(1212553.344455667533))
	fmt.Println(sayBool(true))
	fmt.Println(sayBigInt(400040404))
//	fmt.Println(sayIntArray({1,2,3,4}))
}
