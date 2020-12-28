package models

import (
  u "../utils"
)

func Signup_m(mail string, login string, pass string, id string, vhash string) int {
  db = Db_m()
  rows, junk := db.Query("SELECT COUNT(*) FROM users WHERE email=$1", mail)
  unique := u.CheckCountUnique(rows)
  if (unique != 200) {
    return (unique)
  }
  res, err := db.Exec("INSERT INTO users(id, username, password, email, verify_hash) VALUES ($1, $2, $3, $4, $5);", id, login, pass, mail, vhash)
  return (u.CheckErr(err, 200))

  // gneu gneu declared mais pas identified
  _ = junk
  _ = res
  return (418)
}
