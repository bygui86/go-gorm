package sqlite

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
)

const (
	fileDsnFormat = "%s.db"
	inMemoryDsn   = "file::memory:?cache=shared"
)

func OpenSqliteConnection(dbName string) (*gorm.DB, error) {
	cfg := loadConfig()

	db, openErr := gorm.Open(
		sqlite.Open(buildDsn(cfg, dbName)),
		&gorm.Config{},
	)
	if openErr != nil {
		return nil, openErr
	}

	useErr := db.Use(
		prometheus.New(
			prometheus.Config{
				DBName:          dbName, // use `DBName` as metrics label
				RefreshInterval: 15,     // Refresh metrics interval (default 15 seconds)
				StartServer:     true,   // start http server to expose metrics
				// configure http server port, default port 8080
				// (if you have configured multiple instances, only the first `HTTPServerPort` will be used to start server)
				HTTPServerPort: 9090,
				/*
					user defined metrics implementing

					type MetricsCollector interface {
						Metrics(*Prometheus) []prometheus.Collector
					}
				*/
			},
		),
	)

	return db, useErr
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
