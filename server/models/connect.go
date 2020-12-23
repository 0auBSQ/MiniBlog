package models

import (
  //psql "github.com/jackc/pgx"
  _ "github.com/lib/pq"
  "database/sql"
  "fmt"
)

const (
  host = "localhost"
  port = 5432 // Default postgres port
  user = "postgres" // Default psql user
  password = "test"
  dbname = "mini_blog"
)

var (
  db *sql.DB
)

func Connect_m() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
  db_tmp, err := sql.Open("postgres", psqlInfo)
  if (err != nil) {
    fmt.Printf("Error : %v\n", err)
    panic(err)
  } else {
    fmt.Printf("Db opened without errors.\n")
  }
  db = db_tmp
  err = db.Ping()
  if (err != nil) {
    fmt.Printf("Error : %v\n", err)
    panic(err)
  } else {
    fmt.Printf("Db pinged without errors.\n")
  }
}

func Db_m() *sql.DB {
  return db;
}

func Close_m() error {
  return db.Close()
}
