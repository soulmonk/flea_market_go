package models

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Note struct {
	ID          uint64
	Title       string
	Description string
	CreatedAt   time.Time `bson:"created_at,omitempty" json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty" db:"updated_at"`
}

type Notes struct {
	Notes []Note
}

type NoteDao struct {
	db *sqlx.DB
}

func CreateNoteDao(db *sqlx.DB) *NoteDao {
	dao := NoteDao{db}
	return &dao
}

func (dao *NoteDao) Create(note *Note) error {
	createNoteQuery := `INSERT INTO notes (title, description, created_at, updated_at) VALUES ($1, $2, now(), now()) RETURNING id, created_at, updated_at`
	err := dao.db.QueryRow(createNoteQuery, note.Title, note.Description).Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		log.Println("Error on create note")
		return err
	}
	return nil
}

func (dao *NoteDao) GetAll() (Notes, error) {
	res := Notes{}
	var err error

	rows, err := dao.db.Queryx(`
		SELECT
			id,
			title,
			description,
			created_at,
			updated_at
		FROM Notes
		ORDER BY updated_at DESC`)

	if err != nil {
		log.Println("Error on executing query")
		return res, err
	}

	defer rows.Close()
	for rows.Next() {
		note := Note{}
		err = rows.StructScan(&note)

		if err != nil {
			log.Println("Error corrupted while scanning note")
			return res, err
		}
		log.Println("Fetched note", note)
		res.Notes = append(res.Notes, note)
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error on fetching rows")
		return res, err
	}
	return res, err
}

func (dao *NoteDao) FindById(id string) (Note, error) {
	query := `SELECT * FROM notes where id = $1`
	var note = Note{}

	err := dao.db.Get(&note, query, id)

	if err != nil {
		log.Println("Error on fetching note", err.Error())
		return note, err
	}
	return note, nil
}
