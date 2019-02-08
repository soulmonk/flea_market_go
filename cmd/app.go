package main

import (
	"encoding/json"
	. "first-steps/config"
	"first-steps/pkg"
	models2 "first-steps/pkg/pg/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var config = Config{}
var pgDao = pkg.PGDao{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file and establish a connection to DB
func init() {
	config.Read()

	pgDao.Config = config.Pg
	pgDao.InitDb()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/notes", AllNotesEndPoint).Methods("GET")
	r.HandleFunc("/notes/{id}", FindNoteEndPoint).Methods("GET")
	r.HandleFunc("/notes", CreateNotesEndPoint).Methods("POST")

	addr := ":3000"
	log.Println("listen on", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}

func FindNoteEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	log.Println("Start fetch note with id:", params["id"])
	// TODO find a way convert to int
	note, err := pgDao.NoteDao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Note ID")
		return
	}
	respondWithJson(w, http.StatusOK, note)
}

func CreateNotesEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Start create note")
	var note models2.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		log.Println("Error on decode:", err.Error())
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	log.Println("Parsed Note", note)

	if err := pgDao.NoteDao.Create(&note); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Println("Note created with id", note.ID)

	respondWithJson(w, http.StatusOK, note)
}

func AllNotesEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Start fetch notes")
	notes, err := pgDao.NoteDao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, notes)
}
