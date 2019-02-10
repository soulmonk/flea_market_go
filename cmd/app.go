package main

import (
	"first-steps/config"
	"first-steps/pkg"
	"first-steps/pkg/controllers"
	"first-steps/pkg/db/pg"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var app = pkg.Application{}

// Parse the configuration file and establish a connection to DB
func init() {
	app.Config = config.Load()
	app.PgDao = pg.GetDao(&app.Config.Pg)
}

func main() {
	r := mux.NewRouter()

	controllers.Init(&app, r)

	addr := ":3000"
	log.Println("listen on", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
