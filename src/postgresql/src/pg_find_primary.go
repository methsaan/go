package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"

	"database/sql"
	_ "github.com/lib/pq"
)

func main() {

if len(os.Args) != 6 {
	fmt.Println("usage:", os.Args[0], "HOST", "PORT", "USER", "PASSWORD", "DBNAME")
	return
}

host := os.Args[1]
port0 := os.Args[2]
user := os.Args[3]
password := os.Args[4]
dbname := os.Args[5]

port,err := strconv.Atoi(port0)
if err != nil {
	fmt.Println("postgres: port is not an integer")
}

psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

db, err := sql.Open("postgres", psqlInfo)
if err != nil {
	fmt.Println("postgres: connect error")
}
defer db.Close()

if err != nil {
	fmt.Println("postgres: cannot connect to the database")
}

err = db.Ping()
if err != nil {
	fmt.Println("postgres: cannot reach the database")
}

sqlStatement := `select pg_is_in_recovery();`
var id = "null"

err = db.QueryRow(sqlStatement).Scan(&id)
if err != nil {
	fmt.Println("postgres: select statement in error")
}

switch id {
        case "false":
                fmt.Println(id)
        default:
                panic(errors.New("postgres: no database, exiting"))
                fmt.Println("err")
}

}
