package utils

import (
  "math/rand"
  "time"
  "github.com/gillesdemey/ponysay/ponies"
)

func ListPonies() ([]string) {
  return ponies.AssetNames()
}

func GrabRandomPony() (filePath string) {

  files := ListPonies()

  // Seed the random number generator
  rand.Seed( time.Now().UTC().UnixNano())

  random := rand.Intn(len(files))

  return files[random]
}

func GrabPony(name string) string {
  return "ponies/" + name + ".pony"
}
