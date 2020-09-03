package sqlite

import (
	"github.com/bygui86/go-gorm/utils"
)

const (
	sqliteStorageTypeEnv = "SQLITE_STORAGE_TYPE"

	sqliteStorageTypeDefault = inMemoryStorage
)

func loadConfig() *internalConfig {
	sqliteType := utils.GetStringEnv(sqliteStorageTypeEnv, sqliteStorageTypeDefault.string())

	cfg := &internalConfig{
		storageType: storageType(sqliteType),
	}

	return cfg
}
