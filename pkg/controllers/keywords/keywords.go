package keywords

import (
	"encoding/json"
	"first-steps/pkg"
	"first-steps/pkg/controllers/response"
	"first-steps/pkg/db/pg/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Controller struct {
	dao *models.KeywordDao
}

// TODO
//func logger(v ...interface{}) {
//  var prefix = []string{"Keywords"}
//  log.Println(append(prefix, v...))
//}

func Init(app *pkg.Application, r *mux.Router) {
	//r := mux.NewRouter()
	ctrl := Controller{app.PgDao.KeywordsDao}

	r.HandleFunc("/api/keywords", ctrl.list).Methods("GET")
	r.HandleFunc("/api/keywords/{id}", ctrl.get).Methods("GET")
	r.HandleFunc("/api/keywords", ctrl.create).Methods("POST")
	r.HandleFunc("/api/keywords/{id}", ctrl.remove).Methods("DELETE")

	//return r
}

func (ctrl *Controller) get(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	log.Println("Start fetch model with id:", params["id"])
	// TODO find a way convert to int
	model, err := ctrl.dao.FindById(params["id"])
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, "Invalid Note ID")
		return
	}
	response.RespondWithJson(w, http.StatusOK, model)
}

func (ctrl *Controller) create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Start create model")
	var model models.Keyword
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		log.Println("Error on decode:", err.Error())
		response.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := ctrl.dao.Create(&model); err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println("Model created with id", model.ID)

	response.RespondWithJson(w, http.StatusOK, model)
}

func (ctrl *Controller) list(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Start fetch models")
	models, err := ctrl.dao.List()
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.RespondWithJson(w, http.StatusOK, models)
}

func (ctrl *Controller) remove(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	log.Println("Start removing model with id:", params["id"])
	err := ctrl.dao.Delete(params["id"])
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, "Invalid model ID")
		return
	}
	response.RespondWithJson(w, http.StatusOK, nil)
}
