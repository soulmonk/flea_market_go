package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Note struct {
	ID          uint64
	Title       string
	Description string
	Text        string
	Keywords    []uint64
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

func (dao *NoteDao) Update(note *Note) error {
	return errors.New("not implemented")
}

func (dao *NoteDao) Create(note *Note) error {
	query := `INSERT INTO notes (title, description, "text", created_at, updated_at) VALUES ($1, $2, $3, now(), now()) RETURNING id, created_at, updated_at`
	err := dao.db.
		QueryRow(query, note.Title, note.Description, note.Text).
		Scan(&note.ID, &note.CreatedAt, &note.UpdatedAt)

	if err != nil {
		log.Println("Error on create note")
		return err
	}

	// TODO create relation with keywords

	return nil
}

func (dao *NoteDao) List() (Notes, error) {
	res := Notes{}
	var err error

	// todo with keywords
	rows, err := dao.db.Queryx(`SELECT * FROM notes ORDER BY updated_at DESC`)

	if err != nil {
		log.Println("Error on executing query")
		return res, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Error corrupted while closing rows:", err.Error())
		}
	}()

	for rows.Next() {
		note := Note{}
		if err := rows.StructScan(&note); err != nil {
			log.Println("Error corrupted while scanning note:", err.Error())
			return res, err
		}

		res.Notes = append(res.Notes, note)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error on fetching rows:", err.Error())
		return res, err
	}
	return res, err
}

func (dao *NoteDao) FindById(id string) (Note, error) {
	query := `SELECT * FROM notes where id = $1`
	var note = Note{}

	if err := dao.db.Get(&note, query, id); err != nil {
		log.Println("Error on fetching note", err.Error())
		return note, err
	}
	return note, nil
}

func (dao *NoteDao) Delete(id string) error {
	query := `DELETE FROM notes WHERE id = $1`
	if _, err := dao.db.Exec(query, id); err != nil {
		log.Println("Error on deleting note", err.Error())
		return err
	}
	return nil
}
