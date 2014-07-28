package utils

import "os"

// Detect the terminal type (vterm, xterm, etc.)
func TerminalType() string {
  return os.Getenv("TERM")
}