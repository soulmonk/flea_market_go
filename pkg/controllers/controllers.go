package controllers

import (
	"first-steps/pkg"
	"first-steps/pkg/controllers/notes"
	"github.com/gorilla/mux"
)

func Init(application *pkg.Application, r *mux.Router) {
	// TODO is authenticated
	//r.Handle("/api", )
	notes.Init(application, r)
	InitKeywordsController(application, r)
}
