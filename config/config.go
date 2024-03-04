package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Inti() error {
	return nil
}

func Getlogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
