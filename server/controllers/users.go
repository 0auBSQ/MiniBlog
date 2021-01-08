package controllers

import (
  u "../utils"
)

func UserRights(token string) (string, bool) {
  _, status := Isauth_c(token)
  if (status == 200) {
    id, _ := u.DesDecrypt(token)
    return id, true
  }
  return "", false
}

func AdminRights(token string) (string, bool) {
  admin, status := Isauth_c(token)
  if (status == 200 && admin == 1) {
    id, _ := u.DesDecrypt(token)
    return id, true
  }
  return "", false
}
