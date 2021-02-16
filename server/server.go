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
  "encoding/json"
)

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
    //cookie, err := r.Cookie("session_token")
    token := r.Header.Get("Authorization")
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
    //cookie, err := r.Cookie("session_token")
    token := r.Header.Get("Authorization")
    if (token == "") {
      w.WriteHeader(401)
    } else {
      status := c.CommentDelete_c(token, r.FormValue("qid"))
      w.WriteHeader(status)
    }
  }).Methods("DELETE")

  r.HandleFunc("/api/comment/create", func(w http.ResponseWriter, r *http.Request) {
    //cookie, err := r.Cookie("session_token")
    token := r.Header.Get("Authorization")
    if (token == "") {
      w.WriteHeader(401)
    } else {
      status := c.CommentCreate_c(token, r.FormValue("aid"), r.FormValue("content"))
      w.WriteHeader(status)
    }
  }).Methods("POST")

  r.HandleFunc("/api/comment/update", func(w http.ResponseWriter, r *http.Request) {
    //cookie, err := r.Cookie("session_token")
    token := r.Header.Get("Authorization")
    if (token == "") {
      w.WriteHeader(401)
    } else {
      status := c.CommentUpdate_c(token, r.FormValue("qid"), r.FormValue("content"))
      w.WriteHeader(status)
    }
  }).Methods("PATCH")

  // Articles Routes

  r.HandleFunc("/api/article/read/{aid}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    //cookie, err := r.Cookie("session_token")
    token := r.Header.Get("Authorization")
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
    arts, status := c.ArticleFetch_c(r.FormValue("search"))
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
    //cookie, err := r.Cookie("session_token")
    token := r.Header.Get("Authorization")
    if (token == "") {
      w.WriteHeader(401)
    } else {
      _, status := c.ArticleDelete_c(r.FormValue("aid"), token)
      w.WriteHeader(status)
    }
  }).Methods("DELETE")

  r.HandleFunc("/api/article/create", func(w http.ResponseWriter, r *http.Request) {
    //cookie, err := r.Cookie("session_token")
    token := r.Header.Get("Authorization")
    if (token == "") {
      w.WriteHeader(401)
    } else {
      log.Println(token)
      log.Println("First step");
      aid, status := c.ArticleCreate_c(token, r.FormValue("title"), r.FormValue("content"), r.FormValue("img_link"))
      log.Println(aid);
      log.Println(status);
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
    //cookie, err := r.Cookie("session_token")
    token := r.Header.Get("Authorization")
    if (token == "") {
      w.WriteHeader(401)
    } else {
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
  }).Methods("POST")

  r.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
    token, status, _ := c.Signin_c(r.FormValue("email"), r.FormValue("password"))
    if (status == 200) {
      //expiration := time.Now().Add(3 * 24 * time.Hour)
      //cookie := http.Cookie{Name: "session_token", Value: token, Expires: expiration, HttpOnly: true}
      //cookie_admin := http.Cookie{Name: "status", Value: strconv.Itoa(admin), Expires: expiration}
      //http.SetCookie(w, &cookie)
      //http.SetCookie(w, &cookie_admin)

      w.Header().Set("Authorization", token);
      status = 201; // 201 CREATED
    }
    w.WriteHeader(status)
  }).Methods("POST")

  // Useless while using the local storage
  r.HandleFunc("/api/logout", func(w http.ResponseWriter, r *http.Request) {
    // Delete session cookie by putting a past date
    expiration := time.Time{}
    cookie := http.Cookie{Name: "session_token", Value: "", Expires: expiration, HttpOnly: true}
    http.SetCookie(w, &cookie)
    w.WriteHeader(200)
  }).Methods("GET")

  r.HandleFunc("/api/is_auth/{type}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    //cookie, err := r.Cookie("session_token")
    token := r.Header.Get("Authorization")
    log.Println(token)
    if (token == "") {
      w.WriteHeader(401)
      io.WriteString(w, "No session token")
    } else {
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

  // To fix later
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

  cors := handlers.CORS(
    handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Accept"}),
    handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "PATCH", "DELETE", "OPTIONS"}),
    handlers.ExposedHeaders([]string{"*"}),
    handlers.AllowedOrigins([]string{"http://localhost:8080"}),
    handlers.AllowCredentials(),
  )

  log.Fatal(http.ListenAndServe(":8888", cors(r)))

}
