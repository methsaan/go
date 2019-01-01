package main

import (
	"fmt"
)

func add(f1,f2 float32) float32 {
	return f1+f2
}

func main() {
	var x,y float32 = 4.1,5.4
	fmt.Println("Adding gives",x,y,add(x,y))
}
