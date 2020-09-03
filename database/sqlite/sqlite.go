package sqlite

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	fileDsnFormat = "%s.db"
	inMemoryDsn   = "file::memory:?cache=shared"
)

func OpenSqliteConnection(dbName string) (*gorm.DB, error) {
	cfg := loadConfig()

	return gorm.Open(
		sqlite.Open(buildDsn(cfg, dbName)),
		&gorm.Config{},
	)
}

/*
	dsnExample := "test.db"

	NOTES:
		- You can also use "file::memory:?cache=shared" instead of a path to a file.
		  This will tell SQLite to use a temporary database in system memory.
		  See https://www.sqlite.org/inmemorydb.html
*/
func buildDsn(sqliteCfg *internalConfig, dbName string) string {
	switch sqliteCfg.storageType {
	case fileStorage:
		return fmt.Sprintf(fileDsnFormat, dbName)
	case inMemoryStorage:
		return inMemoryDsn
	default:
		return inMemoryDsn
	}
}
