package main

import (
	"fmt"
	"os"
	"time"
	"strconv"

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

current := time.Now()

port,err := strconv.Atoi(port0)
if err != nil {
	fmt.Println("postgres: port is not an integer", current.String())
	panic(err)
}

psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

t := time.Now()
fmt.Println(t.Format(time.ANSIC))

db, err := sql.Open("postgres", psqlInfo)
if err != nil {
	fmt.Println("postgres: connect error", current.String())
	panic(err)
}
defer db.Close()

if err != nil {
	fmt.Println("postgres: cannot connect to the database", current.String())
	panic(err)
}

err = db.Ping()
if err != nil {
	fmt.Println("postgres: cannot reach the database", current.String())
	panic(err)
}

sqlStatement := `SELECT EXTRACT(MINUTE FROM TIMESTAMP '1970-01-01 00:40:00');`
var id = "null"
err = db.QueryRow(sqlStatement).Scan(&id)
if err != nil {
	fmt.Println("postgres: select statement in error", current.String())
	panic(err)
}

switch id {
	case "40":
		fmt.Println("postgres: connected to database. ip,name,user:", host,dbname,user)
	default:
		fmt.Println("posgres: database error, exiting",id,current.String())
}

}
