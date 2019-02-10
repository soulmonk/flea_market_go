package pg

import (
	"errors"
	"first-steps/config"
	"first-steps/pkg/db/pg/models"
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Dao struct {
	NoteDao *models.NoteDao
	config  *config.PG
}

func GetDao(config *config.PG) *Dao {
	dao := Dao{config: config}
	dao.initConnection()
	dao.initModels()

	return &dao
}

var db *sqlx.DB

func (pg *Dao) Close() error {
	return db.Close()
}

func (pg *Dao) initConnection() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pg.config.Host, pg.config.Port,
		pg.config.User, pg.config.Password, pg.config.Dbname)

	db, err = sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func (pg *Dao) initModels() {
	pg.NoteDao = models.CreateNoteDao(db)
}

func (pg *Dao) Query(sql string) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (pg *Dao) List(sql string) {

}

func (pg *Dao) Insert(sql string) {

}
