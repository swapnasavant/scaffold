package main

import (
  "github.com/codegangsta/negroni"
)

func newLogger() *negroni.Logger {
  l := negroni.NewLogger()
  return l
}
