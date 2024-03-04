package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Inti() error {
	var err error

	// Initialize SQLite
	db, err = initializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func Getlogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
