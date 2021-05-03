package infrastructure

import (
	"database/sql"
	"fmt"
	"gopex/infrastructure/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ===== [Public Functions] ==========
// 指定したDialectorのDB接続情報用インスタンスを生成する
func NewDB(config *config.Config, dialector gorm.Dialector) (*gorm.DB, error) {
	// Create DB Instance and set gorm configs
	db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("[Fatal] Failed to open Mysql Connection Open: %s \n", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("[Fatal] Failed to create database/sql instance: %s \n", err)
	}

	// Set sql.db parameters
	setSqlDBOptions(sqlDB)

	// return initialized db ref
	return db, err
}

// MySQL用のDB接続
func OpenMySQLDatabase(config *config.Config) gorm.Dialector {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Database)

	return mysql.Open(dsn)
}

// ===== [Private Functions] ==========

// sql.db側の制御オプションを設定する
func setSqlDBOptions(sqlDB *sql.DB) {
	// 数値に意図はなく、なんとなくで設定
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
