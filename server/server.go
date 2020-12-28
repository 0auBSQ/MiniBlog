package main

import (
  //psql "github.com/jackc/pgx"
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
  "log"
  "time"
  c "./controllers"
  m "./models"
  //"io"
  "encoding/json"
)

type Register struct {
  Email string
  Username string
  Password string
}

func main() {
  r := mux.NewRouter()

  fmt.Printf("Connecting to database...\n")
  m.Connect_m()
  fmt.Printf("Database connection succeded\n")

  fmt.Printf("Server setup...\n")
  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello world")
  }).Methods("GET")

  r.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)

    info := Register{"test@test.acc", "test", vars["name"]}
    js, err := json.Marshal(info)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)

    fmt.Fprintf(w, "\n%v", c.Signup_c("test@test.acc", "test", vars["name"]))
    //io.WriteString(w, "LIDL")
  }).Methods("GET")

  srv := &http.Server{
    Handler: r,
    Addr: "127.0.0.1:8888",
    WriteTimeout: 15 * time.Second,
    ReadTimeout: 15 * time.Second,
  }

  fmt.Printf("Launched on port 8888\n")
  log.Fatal(srv.ListenAndServe())
}
