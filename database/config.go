package database

import (
	"errors"
	"github.com/bygui86/go-gorm/utils"
)

const (
	typeEnv   = "DB_TYPE"
	dbNameEnv = "DB_NAME"

	typeDefault   = sqliteDb
	dbNameDefault = ""
)

func loadConfig() (*config, error) {
	dbTypeStr := utils.GetStringEnv(typeEnv, typeDefault.string())

	cfg := &config{
		dbType: dbType(dbTypeStr),
		dbName: utils.GetStringEnv(dbNameEnv, dbNameDefault),
	}

	cfgErr := cfg.validateConfig()
	if cfgErr != nil {
		return nil, cfgErr
	}

	return cfg, nil
}

func (c *config) validateConfig() error {
	if c.dbName == dbNameDefault {
		return errors.New("database name must be specified")
	}
	return nil
}
