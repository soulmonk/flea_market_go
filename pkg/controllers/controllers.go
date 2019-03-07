package controllers

import (
	"first-steps/pkg"
	"first-steps/pkg/controllers/keywords"
	"first-steps/pkg/controllers/notes"
	"first-steps/pkg/controllers/user"
	"github.com/gorilla/mux"
)

func Init(application *pkg.Application, r *mux.Router) {

	// TODO mux middleware

	// TODO is authenticated

	user.InitAuth(application, r)
	user.InitUser(application, r)

	//r.Handle("/api", )
	notes.Init(application, r)
	keywords.Init(application, r)
}
