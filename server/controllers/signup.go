package controllers

import (
  "github.com/jzelinskie/whirlpool"
  "encoding/hex"
)

func Signup_c(mail string, login string, pass string) string {
  w := whirlpool.New()
  b_arr := []byte(pass)
  w.Write(b_arr)
  // Placeholder signup func, don't do anything for now
  return (mail + "|" + login + "|" + hex.EncodeToString(w.Sum(nil)))
}
