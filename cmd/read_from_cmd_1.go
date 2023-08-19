package main

import (
	"flag"
	"fmt"
)

func main() {
	var ip = flag.Int("flagname", 1234, "help message")
	fmt.Println("ip has value", *ip)
}
