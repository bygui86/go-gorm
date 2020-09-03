package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dsnFormat = "user=%s password=%s host=%s port=%d %s"
	// if database already exists
	// dsnFormat = "user=%s password=%s host=%s port=%d dbname=%s %s"
)

func OpenPostgresConnection(dbName string) (*gorm.DB, error) {
	cfg, cfgErr := loadConfig()
	if cfgErr != nil {
		return nil, cfgErr
	}

	return gorm.Open(
		postgres.Open(buildDsn(cfg, dbName)),
		&gorm.Config{},
	)
}

/*
	dsnExample := "user=gorm password=gorm dbname=gorm host=127.0.0.1 port=9920 sslmode=disable TimeZone=Europe/Rome"

	NOTES:
		- We are using pgx as postgres’s database/sql driver, it enables prepared statement cache by default
*/
func buildDsn(cfg *internalConfig, dbName string) string {
	return fmt.Sprintf(dsnFormat,
		cfg.username, cfg.password, cfg.host, cfg.port, cfg.params)
	// if database already exists
	//return fmt.Sprintf(dsnFormat,
	//	cfg.username, cfg.password, cfg.host, cfg.port, dbName, cfg.params)
}
