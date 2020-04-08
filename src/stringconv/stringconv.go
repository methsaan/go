package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	s := os.Args[1]
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("s is not an integer")
	}
	fmt.Println(i + 11)
}
