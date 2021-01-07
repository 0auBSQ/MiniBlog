package controllers

import (
  "math/rand"
  u "../utils"
  m "../models"
)

func Signup_c(mail string, login string, pass string) int {
  // Whirlpool hashed (and salted) password (Very strong)
  hashed_pass := u.WhirlpoolHash(pass)

  // Generate strong random user id (using MD5 hash algoritm since it's less sensitive)
  token := make([]byte, 64)
  rand.Read(token)
  unique_id := u.MD5Hash(token)

  // Generate a second hash for account validation
  token = make([]byte, 64)
  rand.Read(token)
  verify_hash := u.MD5Hash(token)

  // Validators here

  // Add user to database
  return (m.Signup_m(mail, login, hashed_pass[0:128], unique_id[0:32], verify_hash[0:32]))
}
