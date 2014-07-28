package main

import (
  "fmt"
  "github.com/gillesdemey/ponysay/utils"
)

// this is where we bootstrap the CLI program

func main() {

  CheckSupport()

  pony := utils.ParsePonyFile(utils.GrabRandomPony())

  fmt.Println(pony.Body)
}

func CheckSupport() {

  term := utils.TerminalType()

  if term == "" || term != "xterm-256color" {
    fmt.Println("terminal type not supported", term)
  }

}
