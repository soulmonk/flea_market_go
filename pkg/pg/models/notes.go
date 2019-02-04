package models

import (
	"database/sql"
	"time"
)

type Note struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type repositories struct {
	Repositories []Note
}

type NoteDao struct {
	db *sql.DB
}

func (dao *NoteDao) getAll() (repositories, error) {
	res := repositories{}
	var err error

	rows, err := dao.db.Query(`
		SELECT
			id,
			repository_owner,
			repository_name,
			total_stars
		FROM repositories
		ORDER BY total_stars DESC`)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		repo := Note{}
		err = rows.Scan(
			&repo.ID,
			&repo.Title,
			&repo.Description,
			&repo.CreatedAt,
			&repo.UpdatedAt,
		)
		if err != nil {
			return res, err
		}
		res.Repositories = append(res.Repositories, repo)
	}
	err = rows.Err()
	if err != nil {
		return res, err
	}
	return res, err
}
