package pkg

import (
	"database/sql"
	"errors"
	"first-steps/config"
	"first-steps/pkg/pg/models"
	"fmt"

	_ "github.com/lib/pq"
)

type PGDao struct {
	Config  config.PG
	NoteDao models.NoteDao
}

var db *sql.DB

func (pg *PGDao) InitDb() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pg.Config.Host, pg.Config.Port,
		pg.Config.User, pg.Config.Password, pg.Config.Dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	initModels(pg)
	fmt.Println("Successfully connected!")
}

func initModels(pg PGDao) {
	pg.NoteDao = models.NoteDao{db}
}

func (pg *PGDao) Query(sql string) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (pg *PGDao) List(sql string) {

}

func (pg *PGDao) Insert(sql string) {

}
