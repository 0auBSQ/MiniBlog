package main

import (
  //psql "github.com/jackc/pgx"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "net/http"
  "fmt"
  "log"
  "time"
  c "./controllers"
  m "./models"
  u "./utils"
  "io"
  "strconv"
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

  // Auth Routes

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
    token, status, admin := c.Signin_c(r.FormValue("email"), r.FormValue("password"))
    if (status == 200) {
      expiration := time.Now().Add(3 * 24 * time.Hour)
      cookie := http.Cookie{Name: "session_token", Value: token, Expires: expiration, HttpOnly: true}
      cookie_admin := http.Cookie{Name: "status", Value: strconv.Itoa(admin), Expires: expiration}
      http.SetCookie(w, &cookie)
      http.SetCookie(w, &cookie_admin)
    }
    w.WriteHeader(status)
  }).Methods("POST")

  r.HandleFunc("/api/logout", func(w http.ResponseWriter, r *http.Request) {
    // Delete session cookie by putting a past date
    expiration := time.Now().Add(-3 * 24 * time.Hour)
    cookie := http.Cookie{Name: "session_token", Value: "", Expires: expiration, HttpOnly: true}
    cookie_admin := http.Cookie{Name: "status", Value: "", Expires: expiration}
    http.SetCookie(w, &cookie)
    http.SetCookie(w, &cookie_admin)
    w.WriteHeader(200)
  }).Methods("GET")

  r.HandleFunc("/api/is_auth", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cookie, err := r.Cookie("session_token")
    if (err != nil) {
      w.WriteHeader(401)
      io.WriteString(w, "No session token")
    } else {
      token := cookie.Value
      admin, status := c.Isauth_c(token)
      if (vars["request_admin_access"] == "yes" && admin == 0) {
        status = 401
      }
      w.WriteHeader(status)
      if (status == 200) {
        io.WriteString(w, "Authorized")
      } else {
        io.WriteString(w, "Invalid token")
      }
    }
  }).Methods("GET")

  fmt.Printf("Launched on port 8888\n")
  log.Fatal(http.ListenAndServe(":8888", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))

}
