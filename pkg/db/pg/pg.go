package pg

import (
	"first-steps/config"
	"first-steps/pkg/db/pg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"

	_ "github.com/lib/pq"
)

type Dao struct {
	NoteDao     *models.NoteDao
	KeywordsDao *models.KeywordDao
	db          *sqlx.DB
}

func GetDao(config *config.PG) *Dao {
	dao := Dao{}
	dao.initConnection(config)
	dao.initModels()

	return &dao
}

func (pg *Dao) Close() error {
	return pg.db.Close()
}

func (pg *Dao) initConnection(config *config.PG) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port,
		config.User, config.Password, config.Dbname)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	pg.db = db

	fmt.Println("Successfully connected!")
}

func (pg *Dao) initModels() {
	pg.NoteDao = models.CreateNoteDao(pg.db)
	pg.KeywordsDao = models.CreateKeywordsDao(pg.db)
}

// TODO not used circular because dependency
func (pg *Dao) Delete(from string, id string, modelName string) error {
	query := `DELETE FROM ` + from + ` WHERE id = $1`
	if _, err := pg.db.Exec(query, id); err != nil {
		log.Println("Error on deleting "+modelName, err.Error())
		return err
	}
	return nil
}

// TODO not used circular because dependency
func (pg *Dao) FindMyId(from string, id string, model *interface{}, modelName string) error {
	query := `DELETE FROM ` + from + ` WHERE id = $1`
	if err := pg.db.Get(model, query, id); err != nil {
		log.Println("Error on fetching "+modelName, err.Error())
		return err
	}
	return nil
}
