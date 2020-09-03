package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dsnFormat = "user=%s password=%s host=%s port=%d dbname=%s %s"
)

func OpenPostgresConnection() (*gorm.DB, error) {
	cfg, cfgErr := loadConfig()
	if cfgErr != nil {
		return nil, cfgErr
	}

	return gorm.Open(
		postgres.Open(buildDsn(cfg)),
		&gorm.Config{},
	)
}

/*
	dsnExample := "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	NOTES:
		- We are using pgx as postgres’s database/sql driver, it enables prepared statement cache by default
*/
func buildDsn(cfg *internalConfig) string {
	return fmt.Sprintf(dsnFormat,
		cfg.username, cfg.password, cfg.host, cfg.port, cfg.dbName, cfg.params)
}
