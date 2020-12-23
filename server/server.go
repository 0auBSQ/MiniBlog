package main

import (
  //psql "github.com/jackc/pgx"
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
  "log"
  "time"
)

func main () {
  r := mux.NewRouter()

  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello world")
  }).Methods("GET")

  r.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello %v", vars["name"])
  }).Methods("GET")

  srv := &http.Server{
    Handler: r,
    Addr: "127.0.0.1:8888",
    WriteTimeout: 15 * time.Second,
    ReadTimeout: 15 * time.Second,
  }

  fmt.Printf("Launched on port 8888")
  log.Fatal(srv.ListenAndServe())
}
