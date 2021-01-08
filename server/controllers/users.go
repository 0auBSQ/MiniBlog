package controllers

import (
  u "../utils"
  m "../models"
)

func UserRights(token string) (string, bool) {
  _, status := Isauth_c(token)
  if (status == 200 && token != "") {
    id, _ := u.DesDecrypt(token)
    return id, true
  }
  return "", false
}

func AdminRights(token string) (string, bool) {
  admin, status := Isauth_c(token)
  if (status == 200 && admin == 1 && token != "") {
    id, _ := u.DesDecrypt(token)
    return id, true
  }
  return "", false
}

func HandleBan_c(token string, uid string) int {
  // Only admins can ban weak chums
  _, admin := AdminRights(token)
  if (admin == false) {
    return 401
  }

  // Drops the ban hammer
  status := m.HandleBan_m(uid)
  return status
}
