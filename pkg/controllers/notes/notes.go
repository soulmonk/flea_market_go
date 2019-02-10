package notes

import (
	"encoding/json"
	"first-steps/pkg"
	"first-steps/pkg/controllers/response"
	"first-steps/pkg/db/pg/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Note struct {
	dao *models.NoteDao
}

func Init(application *pkg.Application, r *mux.Router) {
	//r := mux.NewRouter()
	ctrl := Note{dao: application.PgDao.NoteDao}

	r.HandleFunc("/api/notes", ctrl.list).Methods("GET")
	r.HandleFunc("/api/notes/{id}", ctrl.get).Methods("GET")
	r.HandleFunc("/api/notes", ctrl.create).Methods("POST")
	r.HandleFunc("/api/notes/{id}", ctrl.remove).Methods("DELETE")

	//return r
}

func (ctrl *Note) get(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	log.Println("Start fetch note with id:", params["id"])
	// TODO find a way convert to int
	note, err := ctrl.dao.FindById(params["id"])
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, "Invalid Note ID")
		return
	}
	response.RespondWithJson(w, http.StatusOK, note)
}

func (ctrl *Note) create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Start create note")
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		log.Println("Error on decode:", err.Error())
		response.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := ctrl.dao.Create(&note); err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println("Note created with id", note.ID)

	response.RespondWithJson(w, http.StatusOK, note)
}

func (ctrl *Note) update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("Start update note with id:", params["id"])
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		log.Println("Error on decode:", err.Error())
		response.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
}

func (ctrl *Note) list(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Start fetch notes")
	notes, err := ctrl.dao.List()
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.RespondWithJson(w, http.StatusOK, notes.Notes)
}

func (ctrl *Note) remove(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	log.Println("Start removing note with id:", params["id"])
	// todo find before delete
	err := ctrl.dao.Delete(params["id"])
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, "Invalid Note ID")
		return
	}
	response.RespondWithJson(w, http.StatusOK, nil)
}
