package main

import (
	//"os"
	"os/exec"
	"log"
	"fmt"
	"bytes"
	//"strings"
)

func main() {
    cmd := exec.Command("ls", "-lah")
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }

    outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
    fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)

    //errStr := string(stderr.Bytes())
    //fmt.Printf("err:\n%s\n", errStr)
}
