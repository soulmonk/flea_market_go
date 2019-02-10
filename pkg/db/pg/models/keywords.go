package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
)

type Keyword struct {
	ID   int64
	Name string
}

type KeywordDao struct {
	db *sqlx.DB
}

func CreateKeywordsDao(db *sqlx.DB) *KeywordDao {
	dao := KeywordDao{db}
	return &dao
}

func (dao *KeywordDao) Create(keyword *Keyword) error {
	query := `INSERT INTO keywords (name) VALUES ($1) RETURNING id`

	if err := dao.db.QueryRow(query, keyword.Name).Scan(&keyword.ID); err != nil {
		log.Println("Error on creating keyword: ", err)
		return err
	}

	return nil
}

func (dao *KeywordDao) List() ([]Keyword, error) {
	var res []Keyword
	var err error

	rows, err := dao.db.Queryx(`SELECT * FROM keywords`)

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
		model := Keyword{}
		if err := rows.StructScan(&model); err != nil {
			log.Println("Error corrupted while scanning model:", err.Error())
			return res, err
		}

		res = append(res, model)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error on fetching rows:", err.Error())
		return res, err
	}
	return res, err
}

func (dao *KeywordDao) Update(note *Note) error {
	return errors.New("not implemented")
}

func (dao *KeywordDao) FindById(id string) (Keyword, error) {
	query := `SELECT * FROM keywords where id = $1`
	var keyword = Keyword{}

	if err := dao.db.Get(&keyword, query, id); err != nil {
		log.Println("Error on fetching note", err.Error())
		return keyword, err
	}
	return keyword, nil
}

func (dao *KeywordDao) Delete(id string) error {
	query := `DELETE FROM notes WHERE id = $1`
	if _, err := dao.db.Exec(query, id); err != nil {
		log.Println("Error on deleting model", err.Error())
		return err
	}
	return nil
}
