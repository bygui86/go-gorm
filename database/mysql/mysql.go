package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsnFormat = "%s:%s@tcp(%s:%d)/%s?%s"
)

func OpenMysqlConnection() (*gorm.DB, error) {
	cfg, cfgErr := loadConfig()
	if cfgErr != nil {
		return nil, cfgErr
	}

	return gorm.Open(
		mysql.Open(buildDsn(cfg)),
		&gorm.Config{},
	)
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
func buildDsn(cfg *internalConfig) string {
	return fmt.Sprintf(dsnFormat,
		cfg.username, cfg.password, cfg.host, cfg.port, cfg.dbName, cfg.params)
}
