package main

import (
  "errors"
  "database/sql"
  "fmt"
//  "flag"

  _ "github.com/lib/pq"
)

func main() {

/*
#  h := flag.String("h", "DB host", "String help text")
#  host := *h
#
#  P := flag.Int("P", 5432, "String help text")
#  port := *P
#
#  n := flag.String("n", "DB name", "String help text")
#  dbname := *n
#
#  u := flag.String("u", "DB user", "String help text")
#  user := *u
#
#  p := flag.String("p", "DB pass", "String help text")
#  pass := *p
#
#  flag.Parse()
*/

const (
    host = "127.0.0.1"
    port = 5432
    dbname = "_database"
    user = "_user"
    pass = "welcome"
)

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "pass=%s dbname=%s sslmode=disable",
    host, port, user, pass, dbname)

  db, err := sql.Open("postgres", psqlInfo)

  if err != nil {
    panic(errors.New("postgres: connect error"))
  }
  defer db.Close()

  if err != nil {
    panic(errors.New("Cannot connect to the database"))
  }

fmt.Println("host", host)
fmt.Println("port", port)
fmt.Println("name", dbname)
fmt.Println("user", user)
fmt.Println("pass", pass)

sqlStatement := `SELECT pg_is_in_recovery();`
  var id = "null"
  err = db.QueryRow(sqlStatement).Scan(&id)
  if err != nil {
    panic(errors.New("select statement in error"))
  }

  status := id
switch status {
  case "false":
    fmt.Println("Primary database found at", host)
  default:
    fmt.Println("database error, exiting", id)
}

}
