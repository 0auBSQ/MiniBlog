package utils

import (
  "github.com/jzelinskie/whirlpool"
  "crypto/md5"
  "crypto/aes"
  "crypto/cipher"
  "encoding/hex"
  "math/rand"
  rand2 "crypto/rand"
  "fmt"
  "io"
)

var secretKey []byte;

const (
  salt = "okamari no suzoki kurai sasuke ga kakkoi"
)

func MD5Hash(text []byte) string {
  hash := md5.Sum(text)
  return hex.EncodeToString(hash[:])
}

func WhirlpoolHash(text string) string {
  w := whirlpool.New()
  b_arr := []byte(salt + text)
  w.Write(b_arr)
  return hex.EncodeToString(w.Sum(nil))
}

// Source :
// https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes/
func DesEncrypt(text string) (token string, status int) {
  plain := []byte(text)
  block, err := aes.NewCipher(secretKey)
  if (err != nil) {
    return "", 500
  }
  aesGCM, err := cipher.NewGCM(block)
  if (err != nil) {
    return "", 500
  }
  nonce := make([]byte, aesGCM.NonceSize())
  if _, err = io.ReadFull(rand2.Reader, nonce); err != nil {
    return "", 500
  }
  ciphermsg := aesGCM.Seal(nonce, nonce, plain, nil)
  return fmt.Sprintf("%x", ciphermsg), 200
}

func DesDecrypt(text string) (token string, status int) {
  input, err := hex.DecodeString(text)
  if (err != nil) {
    return "", 401
  }
  block, err := aes.NewCipher(secretKey)
  if (err != nil) {
    return "", 500
  }
  aesGCM, err := cipher.NewGCM(block)
  if (err != nil) {
    return "", 500
  }
  nonceSize := aesGCM.NonceSize()
  nonce, ciphermsg := input[:nonceSize], input[nonceSize:]
  plain, err := aesGCM.Open(nil, nonce, ciphermsg, nil)
  if (err != nil) {
    return "", 401
  }
  return fmt.Sprintf("%s", plain), 200
}

func GenerateRandomKey() {
  // Each time the server is launched, a different random 256-bit AES key is generated (Not optimal, but safe enough for this project)
  secretKey = make([]byte, 32)
  rand.Read(secretKey)
}
