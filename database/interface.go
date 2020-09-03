package database

import (
	"fmt"
	"github.com/bygui86/go-gorm/database/mysql"
	"github.com/bygui86/go-gorm/database/postgres"
	"github.com/bygui86/go-gorm/database/sqlite"
	"github.com/bygui86/go-gorm/model"
	"gopkg.in/logex.v1"
	"gorm.io/gorm"
	"strings"
)

func NewDbInterface() (*DbInterfaceImpl, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	return &DbInterfaceImpl{
		cfg: cfg,
	}, nil
}

func (d *DbInterfaceImpl) OpenConnection() error {
	var db *gorm.DB
	var err error
	switch d.cfg.dbType {
	case sqliteDb:
		db, err = sqlite.OpenSqliteConnection(d.cfg.dbName)
	case postgresDb:
		db, err = postgres.OpenPostgresConnection(d.cfg.dbName)
	case mysqlDb:
		db, err = mysql.OpenMysqlConnection(d.cfg.dbName)
	default:
		err = fmt.Errorf("%s db type not supported", d.cfg.dbType)
	}

	if err != nil {
		return err
	}

	d.db = db

	return nil
}

func (d *DbInterfaceImpl) InitSchema() error {
	logex.Info("Initialize schema")

	switch d.cfg.dbType {
	case postgresDb:
		newErr := d.db.Exec(fmt.Sprintf("CREATE DATABASE %s", d.cfg.dbName)).Error
		if newErr != nil {
			if !strings.Contains(newErr.Error(), fmt.Sprintf("database \"%s\" already exists", d.cfg.dbName)) {
				return newErr
			}
		}
	case mysqlDb:
		newErr := d.db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", d.cfg.dbName)).Error
		if newErr != nil {
			return newErr
		}
		useErr := d.db.Exec(fmt.Sprintf("USE %s;", d.cfg.dbName)).Error
		if useErr != nil {
			return useErr
		}
	}

	migErr := d.db.AutoMigrate(
		&model.Product{},
		&model.Producer{},
	)
	if migErr != nil {
		return migErr
	}

	return nil
}
