package user

import (
	"first-steps/pkg"
	"first-steps/pkg/controllers/response"
	"github.com/gorilla/mux"
	"net/http"
)

type auth struct {
}

func InitAuth(application *pkg.Application, r *mux.Router) {
	ctrl := auth{}

	r.HandleFunc("/api/auth", ctrl.auth).Methods("POST")
}

func (ctrl *auth) auth(w http.ResponseWriter, r *http.Request) {
	response.RespondWithJson(w, http.StatusOK, "auth")
}
