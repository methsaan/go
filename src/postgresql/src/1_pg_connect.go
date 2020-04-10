package main

import (
  "database/sql"
  "fmt"
//  "flag"

  _ "github.com/lib/pq"
)

const (
  host     = "127.0.0.1"
  port     = 5432
  user     = "_owner"
  password = "welcome"
  dbname   = "_database"
)

func main() {

/*
h := flag.String("h", "DB host", "String help text")
host := *h

pt := flag.Int("pt", 5432, "String help text")
port := *pt

n := flag.String("n", "DB name", "String help text")
dbname := *n

u := flag.String("u", "DB user", "String help text")
user := *u

p := flag.String("p", "DB pass", "String help text")
password := *p

flag.Parse()
*/

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
}
