package pkg

import (
	"first-steps/config"
	"first-steps/pkg/db/pg"
)

type Application struct {
	PgDao  *pg.Dao
	Config *config.Config
}
