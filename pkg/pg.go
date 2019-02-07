package pkg

import (
	"errors"
	"first-steps/config"
	"first-steps/pkg/pg/models"
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type PGDao struct {
	Config  config.PG
	NoteDao models.NoteDao
}

var db *sqlx.DB

func (pg *PGDao) InitDb() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pg.Config.Host, pg.Config.Port,
		pg.Config.User, pg.Config.Password, pg.Config.Dbname)

	db, err = sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	initModels(pg)
	fmt.Println("Successfully connected!")
}

func initModels(pg *PGDao) {
	pg.NoteDao = models.NoteDao{}
	pg.NoteDao.SetDb(db)
}

func (pg *PGDao) Query(sql string) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (pg *PGDao) List(sql string) {

}

func (pg *PGDao) Insert(sql string) {

}
