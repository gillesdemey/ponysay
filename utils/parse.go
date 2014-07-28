package utils

import (
  "regexp"
  "github.com/gillesdemey/ponysay/ponies"
)

// Pony object
type Pony struct {
  Body string
}

/**
 * Parse the silly .pony file
 * The syntax doesn't make any sense so regex kung-fu had to be applied
 */
func ParsePonyFile(file string) Pony {

  contents, err := ponies.Asset(file)

  if err != nil {
    panic(err)
  }

  body := extractBody(contents)
  body = replaceBalloonPointers(body, "")
  body = replaceBalloon(body, "")

  pony    := Pony { Body: body }

  return pony
}

// extract the actual pony from the pony file
func extractBody(contents []byte) string {
  return string(regexp.MustCompile(`(?:\n)(?:\${3})([\s\S]*?)$`).
                       FindAllSubmatch(contents, -1)[0][1])

}

// replace $balloon$ something with string
func replaceBalloon(body string, with string) string {
  return regexp.MustCompile(`\$(balloon[0-9]+)\$`).
                ReplaceAllLiteralString(body, with)
}

// replace silly balloon pointers
func replaceBalloonPointers(body string, with string) string {
  return regexp.MustCompile(`(\$\\\$)`).
                ReplaceAllLiteralString(body, with)
}