package controllers

import (
	"encoding/json"
	"first-steps/pkg/db/pg/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func notesController(r *mux.Router) {
	//r := mux.NewRouter()
	r.HandleFunc("/api/notes", list).Methods("GET")
	r.HandleFunc("/api/notes/{id}", get).Methods("GET")
	r.HandleFunc("/api/notes", create).Methods("POST")
	r.HandleFunc("/api/notes/{id}", remove).Methods("DELETE")

	//return r
}

func get(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	log.Println("Start fetch note with id:", params["id"])
	// TODO find a way convert to int
	note, err := app.PgDao.NoteDao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Note ID")
		return
	}
	respondWithJson(w, http.StatusOK, note)
}

func create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Start create note")
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		log.Println("Error on decode:", err.Error())
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := app.PgDao.NoteDao.Create(&note); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println("Note created with id", note.ID)

	respondWithJson(w, http.StatusOK, note)
}

func list(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Start fetch notes")
	notes, err := app.PgDao.NoteDao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, notes.Notes)
}

func remove(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	log.Println("Start removing note with id:", params["id"])
	err := app.PgDao.NoteDao.Delete(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Note ID")
		return
	}
	respondWithJson(w, http.StatusOK, nil)
}
