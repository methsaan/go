package main

import (
	"os"
	"os/exec"
	"fmt"
	"strings"
	"bytes"
	"log"
)

func main() {
	db_host_port_list := os.Getenv("TDCCVT_DB_NODES")
	testArray := strings.Fields(db_host_port_list)
	for _, v := range testArray {
		cmd := exec.Command("/tmp/pg_find_primary", v, "5432", "tdccvt_user", "welcome", "tdccvt")
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}

		//outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
		//fmt.Printf(v,outStr, errStr)

		outStr := string(stdout.Bytes())
		status := strings.TrimSuffix(outStr, "\n")

		fmt.Println(status)
	}
}
