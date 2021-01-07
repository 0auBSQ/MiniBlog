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
  u "./utils"
  //"io"
  //"encoding/json"
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

  u.GenerateRandomKey()

  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello world")
  }).Methods("GET")

  r.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
    /*vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    info := Register{"test@test.acc", "test", vars["name"]}
    js, err := json.Marshal(info)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)*/

    status := c.Signup_c(r.FormValue("email"), r.FormValue("login"), r.FormValue("password"))
    w.WriteHeader(status)
    //fmt.Fprintf(w, "\n%v", c.Signup_c("test@test.acc", "test", vars["name"]))
    //io.WriteString(w, "LIDL")
  }).Methods("POST")

  r.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
    token, status := c.Signin_c(r.FormValue("email"), r.FormValue("password"))
    w.WriteHeader(status)
    if (status == 200) {
      expiration := time.Now().Add(3 * 24 * time.Hour)
      cookie := http.Cookie{Name: "session_token", Value: token, Expires: expiration, HttpOnly: true}
      http.SetCookie(w, &cookie)
    }
  }).Methods("POST")

  srv := &http.Server{
    Handler: r,
    Addr: "127.0.0.1:8888",
    WriteTimeout: 15 * time.Second,
    ReadTimeout: 15 * time.Second,
  }

  fmt.Printf("Launched on port 8888\n")
  log.Fatal(srv.ListenAndServe())
}
