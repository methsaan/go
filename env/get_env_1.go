package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	db_host_port_list := os.Getenv("TDCCVT_DB_NODES")
//        fmt.Println("db nodes: ", db_host_port_list)

	db_ip_port := strings.Split(db_host_port_list, ",")

	fmt.Println("db 1 ip:", db_ip_port[0])
	fmt.Println("db 1 port:", db_ip_port[1])
	fmt.Println("db 2 ip:", db_ip_port[2])
	fmt.Println("db 2 port:", db_ip_port[3])
}
