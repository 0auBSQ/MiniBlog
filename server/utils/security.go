package utils

import (
  "github.com/jzelinskie/whirlpool"
  "crypto/md5"
  "encoding/hex"
)

func MD5Hash(text []byte) string {
  hash := md5.Sum(text)
  return hex.EncodeToString(hash[:])
}

func WhirlpoolHash(text string) string {
  w := whirlpool.New()
  b_arr := []byte(text)
  w.Write(b_arr)
  return hex.EncodeToString(w.Sum(nil))
}
