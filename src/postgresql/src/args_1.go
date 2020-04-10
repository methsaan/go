package main

import (
    "flag"
    "fmt"
//    "strconv"
)

func main() {
    var name string
    flag.StringVar(&name, "opt1", "", "Usage:name")

    var port0 int
    flag.IntVar(&port0, "opt2", 5432, "Usage:port")

    flag.Parse()

/*
    port,err := strconv.Atoi(port0)
    if err != nil {
        fmt.Println("port is not an integer")
        panic(err)
    }
*/

    fmt.Println(name)
    fmt.Println(port0)
}
