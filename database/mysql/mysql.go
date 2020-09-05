package mysql

import (
	"fmt"
	"gopkg.in/logex.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/prometheus"
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
				MetricsCollector: []prometheus.MetricsCollector{
					&prometheus.MySQL{
						//Prefix: "gorm_status_", // Metrics name prefix, default is `gorm_status_`
						//Interval: 15, // Fetch interval, default use Prometheus's RefreshInterval
						//VariableNames: []string{"Threads_running"}, // Select variables from SHOW STATUS, if not set, uses all status variables
					},
				},
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
