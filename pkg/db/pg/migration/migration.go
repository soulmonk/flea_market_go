package migration

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type migrationFile struct {
	path string
	name string
}

func Proceed(db *sqlx.DB) error {
	var err error
	var count uint64
	if count, err = initTable(db); err != nil {
		return err
	}

	log.Println("Existing migration count: ", count)

	var procideMigration = map[string]bool{}
	if err = loadMigration(db, &procideMigration); err != nil {
		return err
	}

	_, err = getMigrationList()
	if err != nil {
		return err
	}

	return err
}

func create(db *sqlx.DB, name string) error {
	query := `INSERT INTO migration (name, created_at) VALUES ($1, now()) RETURNING id`
	var id uint64
	err := db.
		QueryRow(query, name).
		Scan(&id)

	if err != nil {
		log.Println("Error on create note")
		return err
	}
	log.Println("Created ne record:", id)

	return nil
}

func loadMigration(db *sqlx.DB, migrations *map[string]bool) error {
	var err error

	rows, err := db.Queryx(`SELECT name FROM migration`)

	if err != nil {
		log.Println("Error on executing query")
		return err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Error corrupted while closing rows:", err.Error())
		}
	}()

	var name string
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			log.Println("Error corrupted while scanning migration:", err.Error())
			return err
		}

		// TODO
		// TODO
		// TODO
		migrations[name] = true
	}
	if err := rows.Err(); err != nil {
		log.Println("Error on fetching rows:", err.Error())
		return err
	}
	return err
}

func initTable(db *sqlx.DB) (uint64, error) {

	query := `SELECT COUNT(1) FROM migration;`
	var count uint64
	err := db.QueryRow(query).Scan(&count)

	if err != nil {
		// todo check error table dose not exists
		log.Println("Some error on get count from migration", err.Error())
		return count, err
		query = `CREATE TABLE migration(
  id serial,
  name varchar(255),
  created_at timestamp default now()
);

create unique index migration_id_uindex
  on migration (id);

alter table migration
  add constraint migration_pk
    primary key (id);
`

		if _, err := db.Exec(query); err != nil {
			log.Println("Error on deleting note", err.Error())
			return count, err
		}
	}

	return count, nil
}

func getMigrationList() (*[]migrationFile, error) {
	var list []migrationFile

	return &list, nil
}
