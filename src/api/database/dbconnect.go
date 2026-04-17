package database

import (
	"sync"

	"github.com/drumer2142/microWeb/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	rw sync.RWMutex
)

// Init opens a single shared DB pool. Safe to call multiple times after a successful first open.
func Init() error {
	rw.Lock()
	defer rw.Unlock()
	if db != nil {
		return nil
	}
	var err error
	db, err = gorm.Open(mysql.Open(config.DBURL), &gorm.Config{})
	return err
}

// Get returns the shared handle. Init must have succeeded first.
func Get() *gorm.DB {
	rw.RLock()
	defer rw.RUnlock()
	return db
}

// Close releases the underlying pool.
func Close() error {
	rw.Lock()
	defer rw.Unlock()
	if db == nil {
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	db = nil
	return err
}
