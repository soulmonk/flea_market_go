package main

import (
	"first-steps/config"
	"first-steps/pkg/db/pg"
	"first-steps/pkg/db/pg/migration"
	"log"
)

func main() {
	log.Println("Starting migration")
	cfg := config.Load()
	db := pg.InitConnection(&cfg.Pg)

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := migration.Proceed(db); err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
}
