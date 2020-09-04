package postgres

import (
	"fmt"
	"gopkg.in/logex.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

const (
	dsnFormat = "user=%s password=%s host=%s port=%d dbname=%s %s" // if database already exists
)

func OpenPostgresConnection(dbName string) (*gorm.DB, error) {
	cfg, cfgErr := loadConfig()
	if cfgErr != nil {
		return nil, cfgErr
	}

	var db *gorm.DB
	var openErr error
	db, openErr = gorm.Open(
		postgres.Open(buildDsn(cfg, dbName, true)),
		&gorm.Config{},
	)
	if openErr != nil {
		if strings.Contains(openErr.Error(), fmt.Sprintf("Unknown database '%s'", dbName)) {
			logex.Infof("Database %s not yet created, connecting to MySQL without dbName...", dbName)
			db, openErr = gorm.Open(
				postgres.Open(buildDsn(cfg, dbName, false)),
				&gorm.Config{},
			)
		}
	}

	return db, openErr
}

/*
	dsnExample := "user=gorm password=gorm dbname=gorm host=127.0.0.1 port=9920 sslmode=disable TimeZone=Europe/Rome"

	NOTES:
		- We are using pgx as postgres’s database/sql driver, it enables prepared statement cache by default
*/
func buildDsn(cfg *internalConfig, dbName string, useDbName bool) string {
	if useDbName { // if the database already exists
		return fmt.Sprintf(dsnFormat,
			cfg.username, cfg.password, cfg.host, cfg.port, dbName, cfg.params)
	} else {
		return fmt.Sprintf(dsnFormat,
			cfg.username, cfg.password, cfg.host, cfg.port, cfg.params)
	}
}
