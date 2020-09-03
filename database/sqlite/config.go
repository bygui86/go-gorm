package sqlite

import (
	"errors"
	"github.com/bygui86/go-gorm/utils"
)

const (
	sqliteStorageTypeEnv = "SQLITE_STORAGE_TYPE"
	sqliteFilenameEnv    = "SQLITE_FILENAME"

	sqliteStorageTypeDefault = inMemoryStorage
	sqliteFilenameDefault    = ""
)

func loadConfig() (*internalConfig, error) {
	sqliteType := utils.GetStringEnv(sqliteStorageTypeEnv, sqliteStorageTypeDefault.string())

	cfg := &internalConfig{
		storageType: storageType(sqliteType),
		filename:    utils.GetStringEnv(sqliteFilenameEnv, sqliteFilenameDefault),
	}

	cfgErr := cfg.validateConfig()
	if cfgErr != nil {
		return nil, cfgErr
	}

	return cfg, nil
}

func (c *internalConfig) validateConfig() error {
	if c.storageType == fileStorage {
		if c.filename == sqliteFilenameDefault {
			return errors.New("filename must be specified")
		}
	}
	return nil
}
