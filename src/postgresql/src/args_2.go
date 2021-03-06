package main

import (
    "flag"
    "fmt"
    "errors"

    "database/sql"
    _ "github.com/lib/pq"
)

func main() {
    var host string
    flag.StringVar(&host, "host", "", "Usage:host")

    var user string
    flag.StringVar(&user, "user", "", "Usage:user")

    var password string
    flag.StringVar(&password, "password", "", "Usage:password")

    var dbname string
    flag.StringVar(&dbname, "dbname", "", "Usage:dbname")

    var port int
    flag.IntVar(&port, "port", 5432, "Usage:port")

    flag.Parse()

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
