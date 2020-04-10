package main

import (
    "flag"
    "fmt"

    "database/sql"
    _ "github.com/lib/pq"
)

func main() {
    host := flag.String("host", "127.0.0.1", "database ip")
    user := flag.String("user", "user1", "database user")
    dbname := flag.String("dbname", "db1", "database name")
    password := flag.String("password", "welcome", "database password")

    var port int
    flag.IntVar(&port, "port", 5432, "database port")

    flag.Parse()

    fmt.Printf("Value of host is: %s\n", *host)
    fmt.Printf("Value of user is: %s\n", *user)
    fmt.Printf("Value of dbname is: %s\n", *dbname)
    fmt.Printf("Value of password is: %s\n", *password)
    fmt.Printf("Value of port is: %d\n", port)


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
