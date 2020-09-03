package database

import (
	"github.com/bygui86/go-gorm/utils"
)

const (
	typeEnv = "DB_TYPE"

	typeDefault = sqliteDb
)

func loadConfig() *config {
	dbTypeStr := utils.GetStringEnv(typeEnv, typeDefault.string())

	cfg := &config{
		dbType: dbType(dbTypeStr),
	}

	return cfg
}
