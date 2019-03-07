package main

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

var host string
var port int
var user string
var password string
var dbname string

func main() {
fmt.Print("Enter the host: ")
fmt.Scanf("%s", &host)
fmt.Print("Enter the port: ")
fmt.Scanf("%d", &port)
fmt.Print("Enter the user: ")
fmt.Scanf("%s", &user)
fmt.Print("Enter the password: ")
fmt.Scanf("%s", &password)
fmt.Print("Enter the dbname: ")
fmt.Scanf("%s", &dbname)

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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
