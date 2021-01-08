package controllers

import (
  u "../utils"
  m "../models"
)

func Isauth_c(token string) (int, int) {
  if (token == "") {
    return 0, 401
  }
  id, status := u.DesDecrypt(token)
  if (status != 200) {
    return 0, status
  }
  user, status := m.Userinfo_m(id)
  return user.Is_admin, status
}

func Signin_c(mail string, pass string) (token string, status int, admin int) {
  // Hash the password first
  hashed_pass := u.WhirlpoolHash(pass)

  // Check if the credentials return an existing user in the database
  id, status, admin := m.Signin_m(mail, hashed_pass)

  // If success generate a DES encrypted secret token (no JWT)
  if (status == 200) {
    token, status := u.DesEncrypt(id)
    return token, status, admin
  }

  // Else return the error status
  return "", status, 0
}
