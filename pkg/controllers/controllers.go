package controllers

import (
	"first-steps/pkg"
	"github.com/gorilla/mux"
)

var app *pkg.Application

func Init(application *pkg.Application, r *mux.Router) {
	app = application

	r.Handle("/api/", notesController())
}
