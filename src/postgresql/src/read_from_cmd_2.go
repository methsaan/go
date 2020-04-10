package main

import (
    "flag"
    "fmt"
)

func main() {
    host := flag.String("h", "127.0.0.1", "database host ip")

    var port int
    flag.IntVar(&port, "p", 5555, "database port")

    flag.Parse()

    fmt.Printf("Value of host is: %s\n", *host)
    fmt.Printf("Value of port is: %d\n", port)
}
