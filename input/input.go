package main

import(
	"fmt"
)

func main() {
	var name string
	var age int
	fmt.Println("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Enter an age: ")
	fmt.Scanf("%d", &age)
	fmt.Printf("%s is %d years old\n", name, age)
}
