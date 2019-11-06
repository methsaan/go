package main

import (
	"fmt"
)

func sayHello(s string) string {
	return "Hello " + s
}

func main() {
	fmt.Println(sayHello("George"))
	var age int = 12
	fmt.Println(age)
	var smallfloat float32 = 53.123456
	fmt.Println(smallfloat)
	var bigfloat float64 = 1212553.123456789012345
	fmt.Println(bigfloat)
	var booleanVar bool = true
	fmt.Println(booleanVar)
	var bigInt uint = 4000000000
	fmt.Println(bigInt)
	var str string = "Hellos"
	fmt.Println(str)
	var numArr[10] int
	numArr[0] = 1253
	numArr[1] = 14
	numArr[2] = 125
	numArr[3] = 124
	numArr[4] = 1543
	numArr[5] = 723
	numArr[6] = 345
	numArr[7] = 5234
	numArr[8] = 8235
	numArr[9] = 908773
	fmt.Println(numArr)
}
