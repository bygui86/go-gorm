package database

import (
	"fmt"
	"github.com/bygui86/go-gorm/database/mysql"
	"github.com/bygui86/go-gorm/database/postgres"
	"github.com/bygui86/go-gorm/database/sqlite"
	"github.com/bygui86/go-gorm/model"
	"gopkg.in/logex.v1"
	"gorm.io/gorm"
)

func NewDbInterface() (*DbInterfaceImpl, error) {
	cfg := loadConfig()

	return &DbInterfaceImpl{
		cfg: cfg,
	}, nil
}

func (d *DbInterfaceImpl) OpenConnection() error {
	var db *gorm.DB
	var err error
	switch d.cfg.dbType {
	case sqliteDb:
		db, err = sqlite.OpenSqliteConnection()
	case postgresDb:
		db, err = postgres.OpenPostgresConnection()
	case mysqlDb:
		db, err = mysql.OpenMysqlConnection()
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

	err := d.db.AutoMigrate(
		&model.Product{},
		&model.Producer{},
	)
	if err != nil {
		return err
	}

	return nil
}
