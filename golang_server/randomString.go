package main

// Stole this from https://www.calhoun.io/creating-random-strings-in-go/

import (  
  "math/rand"
  "time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +  
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(  
  rand.NewSource(time.Now().UnixNano()))

func randomStringWithCharset(length int, charset string) string {  
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func randomString(length int) string {  
  return randomStringWithCharset(length, charset)
}
