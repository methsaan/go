package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World - output example")
	fmt.Print("Hello World - output w/o newline\t\t")
	fmt.Printf("format %s\n", "WHat ever")
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Welcome to Go Programming " + name)
	fmt.Println("Welcome to Go Programming", name)
	var age float32
	fmt.Print("What is your name and age? ")
	fmt.Scanf("%s is %f", &name, &age)
	fmt.Println(name, "is", age, "years old")
}
