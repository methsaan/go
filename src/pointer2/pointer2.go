package main

import "fmt"

func main(){
	var ptr *float64 = new(float64)
	*ptr = 4.5;
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
