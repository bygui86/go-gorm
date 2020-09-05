package postgres

import (
	"fmt"
	"gopkg.in/logex.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
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
