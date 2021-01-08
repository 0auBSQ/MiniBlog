package models

import (
  u "../utils"
)

type User struct {
  Id string
  Username string
  Email string
  Verified int
  Verify_hash string
  Is_admin int
  Is_banned int
  Last_log string
  Reg_date string
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

func Signin_m(mail string, pass string) (id string, status int, admin int) {
  db = Db_m()
  id = ""
  admin = 0
  rows, err := db.Query("SELECT id, is_admin FROM users WHERE email=$1 AND password=$2", mail, pass)
  if (err != nil) {
    return "", 500, 0
  }
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&id, &admin)
    if (err != nil) {
      return "", 500, 0
    }
  }
  err = rows.Err()
  if (err != nil) {
    return "", 500, 0
  }

  // If the id is not found, its value will remains an empty string, so the user does not exist, unauthorized status
  if (id == "") {
    return "", 404, 0
  }
  return id, 200, admin
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
    err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Verified, &user.Verify_hash, &user.Is_admin, &user.Is_banned, &user.Last_log, &user.Reg_date)
    if (err != nil) {
      return user, 500
    }
    status = 200
  }
  return user, status
}

func HandleBan_m(id string) int {
  db = Db_m()
  _, err := db.Exec("UPDATE users SET is_banned=1 WHERE id=$1;", id)
  return (u.CheckErr(err, 200))
}
