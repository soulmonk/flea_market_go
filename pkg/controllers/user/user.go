package user

import (
	"first-steps/pkg"
	"first-steps/pkg/controllers/response"
	"github.com/gorilla/mux"
	"net/http"
)

type user struct {
}

func InitUser(application *pkg.Application, r *mux.Router) {
	ctrl := user{}

	r.HandleFunc("/api/user/me", ctrl.me).Methods("GET")
}

func (ctrl *user) me(w http.ResponseWriter, r *http.Request) {
	response.RespondWithJson(w, http.StatusOK, "me")
}
