package main

import (
	"os"
	"fmt"

	"strings"
        "strconv"

	"database/sql"
	_ "github.com/lib/pq"
)

func main() {

if len(os.Args) != 5 {
        fmt.Println("usage:", os.Args[0], "PORT", "USER", "PASSWORD", "DBNAME")
        return
}

port0 := os.Args[1]
user := os.Args[2]
password := os.Args[3]
dbname := os.Args[4]

port,err := strconv.Atoi(port0)
if err != nil {
	fmt.Println("postgres: port is not an integer",err)
}

	db_host_port_list := os.Getenv("TDCCVT_DB_NODES")
	testArray := strings.Fields(db_host_port_list)
	for _, host := range testArray {
		status := db_status(host,port,user,password,dbname)
		if status == "false" {
			fmt.Println(host)
		}
	}
}

// function called to check database connectivity,
func db_status(host string, port int, user string, password string, dbname string) string {

psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

db, err := sql.Open("postgres", psqlInfo)
if err != nil {
	fmt.Println("postgres: connect error", err)
}
defer db.Close()

if err != nil {
	fmt.Println("postgres: cannot connect to the database ", err)
}

err = db.Ping()
if err != nil {
	fmt.Println("postgres: cannot reach the database ", err)
}

sqlStatement := `select pg_is_in_recovery();`
var id = "null"

err = db.QueryRow(sqlStatement).Scan(&id)
if err != nil {
	fmt.Println("postgres: select statement in error ", err)
}

return id

}
