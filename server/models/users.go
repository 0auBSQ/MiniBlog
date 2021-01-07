package models

import (
  u "../utils"
)

type User struct {
  id string
  username string
  email string
  verified int
  verify_hash string
  is_admin int
  is_banned int
  last_log string
  reg_date string
}

func Signup_m(mail string, login string, pass string, id string, vhash string) int {
  db = Db_m()
  rows, err := db.Query("SELECT COUNT(*) FROM users WHERE email=$1", mail)
  if (err != nil) {
    return (500)
  }
  unique := u.CheckCountUnique(rows)
  if (unique != 200) {
    return (unique)
  }
  _, err = db.Exec("INSERT INTO users(id, username, password, email, verify_hash) VALUES ($1, $2, $3, $4, $5);", id, login, pass, mail, vhash)
  return (u.CheckErr(err, 200))
}

func Signin_m(mail string, pass string) (id string, status int) {
  db = Db_m()
  id = ""
  rows, err := db.Query("SELECT id FROM users WHERE email=$1 AND password=$2", mail, pass)
  if (err != nil) {
    return "", 500
  }
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&id)
    if (err != nil) {
      return "", 500
    }
  }
  err = rows.Err()
  if (err != nil) {
    return "", 500
  }

  // If the id is not found, its value will remains an empty string, so the user does not exist, unauthorized status
  if (id == "") {
    return "", 401
  }
  return id, 200
}

func Userinfo_m(id string) (user User, status int) {
  db = Db_m()
  status = 400
  rows, err := db.Query("SELECT id, username, email, verified, verify_hash, is_admin, is_banned, last_log, reg_date FROM users WHERE id=$1", id)
  if (err != nil) {
    return user, 500
  }
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&user.id, &user.username, &user.email, &user.verified, &user.verify_hash, &user.is_admin, &user.is_banned, &user.last_log, &user.reg_date)
    if (err != nil) {
      return user, 500
    }
    status = 200
  }
  return user, status
}
