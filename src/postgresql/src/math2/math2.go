package main

import (
	"fmt"
)

func add(x,y float64) float64 {
	return x+y
}

func main() {
	var num1,num2 float64 = 5.566,-91.2343
	//var num1,num2 := 5.566,-91.2343
	fmt.Println("Adding and gives",num1,num2,add(num1,num2))
}
