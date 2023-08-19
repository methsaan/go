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
	fmt.Println("Usage:", os.Args[0], "HOST", "PORT", "USER", "PASSWORD", "DBNAME")
	return
}

host := os.Args[1]
port0 := os.Args[2]
user := os.Args[3]
password := os.Args[4]
dbname := os.Args[5]

port,err := strconv.Atoi(port0)
if err != nil {
	fmt.Println("port is not an integer")
	panic(err)
}

psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

db, err := sql.Open("postgres", psqlInfo)
if err != nil {
	panic(err)
	panic(errors.New("postgres: connect error"))
}
defer db.Close()

if err != nil {
	panic(err)
	panic(errors.New("Cannot connect to the database"))
}

err = db.Ping()
if err != nil {
	panic(err)
	panic(errors.New("cannot reach the database"))
}

sqlStatement := `SELECT pg_is_in_recovery();`
var id = "null"
err = db.QueryRow(sqlStatement).Scan(&id)
if err != nil {
	panic(errors.New("select statement in error"))
}

switch id {
	case "false":
		fmt.Println("Primary database found at", host)
	default:
		fmt.Println("database error, exiting", id)
		panic(errors.New("no primary database, exiting"))
}

}
