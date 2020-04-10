package main

import (
    "errors"
    "fmt"
)

func saveEarth() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = r.(error)
        }
    }()
    TooLate()
    return
}

func TooLate() {
    A()
    panic(errors.New("Then there's nothing we can do"))
}

func A() {
    defer fmt.Println("If it's more than 100 degrees...")
}

func main() {
    err := saveEarth()
    fmt.Println(err)
}
