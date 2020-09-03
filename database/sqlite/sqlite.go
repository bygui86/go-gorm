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

func OpenSqliteConnection() (*gorm.DB, error) {
	cfg, cfgErr := loadConfig()
	if cfgErr != nil {
		return nil, cfgErr
	}

	return gorm.Open(
		sqlite.Open(buildDsn(cfg)),
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
func buildDsn(sqliteCfg *internalConfig) string {
	switch sqliteCfg.storageType {
	case fileStorage:
		return fmt.Sprintf(fileDsnFormat, sqliteCfg.filename)
	case inMemoryStorage:
		return inMemoryDsn
	default:
		return inMemoryDsn
	}
}
