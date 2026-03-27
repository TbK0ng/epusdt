//go:build !sqlite_cgo

package dao

import (
	"github.com/libtnb/sqlite"
	"gorm.io/gorm"
)

func openDB(dsn string, cfg *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dsn), cfg)
}
