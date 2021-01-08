package main

import (
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

  u.GenerateRandomKey()

  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello world")
  }).Methods("GET")

  // Comment routes

  r.HandleFunc("/api/comment/fetch/{aid}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cookie, err := r.Cookie("session_token")
    token := ""
    if (err == nil) {
      token = cookie.Value
    }
    coms, status := c.CommentFetch_c(token, vars["aid"])
    js, err := json.Marshal(coms)
    if (err != nil) {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(js)
  }).Methods("GET")

  r.HandleFunc("/api/comment/delete", func(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session_token")
    if (err != nil) {
      w.WriteHeader(401)
    } else {
      token := cookie.Value
      status := c.CommentDelete_c(token, r.FormValue("qid"))
      w.WriteHeader(status)
    }
  }).Methods("DELETE")

  r.HandleFunc("/api/comment/create", func(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session_token")
    if (err != nil) {
      w.WriteHeader(401)
    } else {
      token := cookie.Value
      status := c.CommentCreate_c(token, r.FormValue("aid"), r.FormValue("content"))
      w.WriteHeader(status)
    }
  }).Methods("POST")

  r.HandleFunc("/api/comment/update", func(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session_token")
    if (err != nil) {
      w.WriteHeader(401)
    } else {
      token := cookie.Value
      status := c.CommentUpdate_c(token, r.FormValue("qid"), r.FormValue("content"))
      w.WriteHeader(status)
    }
  }).Methods("PATCH")

  // Articles Routes

  r.HandleFunc("/api/article/read/{aid}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cookie, err := r.Cookie("session_token")
    token := ""
    if (err == nil) {
      token = cookie.Value
    }
    art, status := c.ArticleRead_c(token, vars["aid"])
    js, err := json.Marshal(art)
    if (err != nil) {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(js)
  }).Methods("GET")

  r.HandleFunc("/api/article/fetch", func(w http.ResponseWriter, r *http.Request) {
    arts, status := c.ArticleFetch_c()
    js, err := json.Marshal(arts)
    if (err != nil) {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(js)
  }).Methods("GET")

  r.HandleFunc("/api/article/delete", func(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session_token")
    if (err != nil) {
      w.WriteHeader(401)
    } else {
      token := cookie.Value
      _, status := c.ArticleDelete_c(r.FormValue("aid"), token)
      w.WriteHeader(status)
    }
  }).Methods("DELETE")

  r.HandleFunc("/api/article/create", func(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session_token")
    if (err != nil) {
      w.WriteHeader(401)
    } else {
      token := cookie.Value
      aid, status := c.ArticleCreate_c(token, r.FormValue("title"), r.FormValue("content"), r.FormValue("img_link"))
      js, err := json.Marshal(aid)
      if (err != nil) {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(status)
      w.Write(js)
    }
  }).Methods("POST")

  r.HandleFunc("/api/article/update", func(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session_token")
    if (err != nil) {
      w.WriteHeader(401)
    } else {
      token := cookie.Value
      aid, status := c.ArticleUpdate_c(r.FormValue("aid"), token, r.FormValue("title"), r.FormValue("content"), r.FormValue("img_link"))
      js, err := json.Marshal(aid)
      if (err != nil) {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(status)
      w.Write(js)
    }
  }).Methods("PATCH")

  // Auth Routes

  r.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
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
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    http.SetCookie(w, &cookie)
    http.SetCookie(w, &cookie_admin)
    w.WriteHeader(200)
  }).Methods("GET")

  r.HandleFunc("/api/is_auth/{type}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cookie, err := r.Cookie("session_token")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    if (err != nil) {
      w.WriteHeader(401)
      io.WriteString(w, "No session token")
    } else {
      token := cookie.Value
      admin, status := c.Isauth_c(token)
      if (vars["type"] == "admin" && admin == 0) {
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

  r.HandleFunc("/api/ban/{uid}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    cookie, err := r.Cookie("session_token")
    if (err != nil) {
      w.WriteHeader(401)
    } else {
      token := cookie.Value
      status := c.HandleBan_c(token, vars["uid"])
      w.WriteHeader(status)
    }
  }).Methods("PATCH")

  fmt.Printf("Launched on port 8888\n")
  log.Fatal(http.ListenAndServe(":8888", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "PATCH", "DELETE", "OPTIONS"}), handlers.AllowedOrigins([]string{"http://localhost:8080"}))(r)))

}
