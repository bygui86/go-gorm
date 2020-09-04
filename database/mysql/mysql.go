package mysql

import (
	"fmt"
	"gopkg.in/logex.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

const (
	dsnFormat         = "%s:%s@tcp(%s:%d)/%s?%s" // if the database already exists
	dsnFormatNoDbName = "%s:%s@tcp(%s:%d)/?%s"
)

func OpenMysqlConnection(dbName string) (*gorm.DB, error) {
	cfg, cfgErr := loadConfig()
	if cfgErr != nil {
		return nil, cfgErr
	}

	var db *gorm.DB
	var openErr error
	db, openErr = gorm.Open(
		mysql.Open(buildDsn(cfg, dbName, true)),
		&gorm.Config{},
	)
	if openErr != nil {
		if strings.Contains(openErr.Error(), fmt.Sprintf("Unknown database '%s'", dbName)) {
			logex.Infof("Database %s not yet created, connecting to MySQL without dbName...", dbName)
			db, openErr = gorm.Open(
				mysql.Open(buildDsn(cfg, dbName, false)),
				&gorm.Config{},
			)
		}
	}

	return db, openErr
}

/*
	refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	dsnExample := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	NOTES:
		- To handle time.Time correctly, you need to include parseTime as a parameter.
		  For more parameters see https://github.com/go-sql-driver/mysql#parameters
		- To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4.
		  For a detailed explanation see https://mathiasbynens.be/notes/mysql-utf8mb4
*/
func buildDsn(cfg *internalConfig, dbName string, useDbName bool) string {
	if useDbName { // if the database already exists
		return fmt.Sprintf(dsnFormat,
			cfg.username, cfg.password, cfg.host, cfg.port, dbName, cfg.params)
	} else {
		return fmt.Sprintf(dsnFormatNoDbName,
			cfg.username, cfg.password, cfg.host, cfg.port, cfg.params)
	}
}
