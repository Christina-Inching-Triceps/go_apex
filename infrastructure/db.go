package infrastructure

import (
	"database/sql"
	"fmt"
	"gopex/infrastructure/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ===== [Public Functions] ==========

func NewDB(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database)

	// Create DB Instance
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("[Fatal] Failed to open Mysql Connection Open: %s \n", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("[Fatal] Failed to create database/sql instance: %s \n", err)
	}

	// Set connection settings
	setupSqlDB(sqlDB)

	// return initialized db ref
	return db, err
}

// ===== [Private Functions] ==========

func setupSqlDB(sqlDB *sql.DB) {
	// TODO: 特に根拠はなくなんとなく設定
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
