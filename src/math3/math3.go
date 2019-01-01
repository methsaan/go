package main

import (
	"fmt"
)

func add(f1,f2 float64) float64 {
	return f1+f2
}

func main() {
	x,y := 34.2,-33.4
	fmt.Println("Adding and gives",x,y,add(x,y))
}
