package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

const (
	OK_RESPONSE = "{\"status\": \"ok\"}"
)

var (
	logger = newLogger()
)

func main() {

	r := mux.NewRouter()

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewStatic(http.Dir("public")))
	n.Use(logger)
	n.UseHandler(r)

	n.Run(":9002")
}
