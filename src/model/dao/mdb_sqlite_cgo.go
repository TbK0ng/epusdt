//go:build sqlite_cgo

package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openDB(dsn string, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dsn), cfg)
}
