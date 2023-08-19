package main

import (
	"fmt"
	"os"
	"strconv"
	"log"

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
	log.Fatal("postgres: illegal port value: ", err)
}

psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

// https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
db, err := sql.Open("postgres", psqlInfo)
if err != nil {
	log.Fatal("postgres: github.com/lib/pq issues:", err)
}
defer db.Close()

err = db.Ping()
if err != nil {
	log.Fatal("postgres: error reported:", err)
}

sqlStatement := `SELECT EXTRACT(MINUTE FROM TIMESTAMP '1970-01-01 00:40:00');`

var id = "null"
err = db.QueryRow(sqlStatement).Scan(&id)
if err != nil {
	log.Fatal("postgres: select statemet in error:", err)
}

switch id {
	case "40":
		fmt.Println("postgres: primary database found: host,status:", host,id)
	default:
		log.Fatal("postgres: no primary database found:", err)
		fmt.Println(host,id)
}

}
