package postgresql

import (
	"fmt"
	"github.com/RichFastTrade/rft_shared/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

var db *gorm.DB
var once sync.Once

func getDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = Connect(conf.Get().Postgresql.Database)
		if err != nil {
			panic(err)
		}
	})
	return db
}

func Connect(dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Get().Postgresql.Host,
		conf.Get().Postgresql.Port,
		conf.Get().Postgresql.User,
		conf.Get().Postgresql.Password,
		dbname,
	)
	connectDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := connectDB.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)

	return connectDB, nil
}

func Create(dbname string) error {
	var exists bool
	if err := getDB().Raw("SELECT EXISTS(SELECT datname FROM pg_database WHERE datname = ?)", dbname).Scan(&exists).Error; err != nil {
		return err
	}
	if !exists {
		if err := getDB().Exec("CREATE DATABASE " + dbname).Error; err != nil {
			return err
		}
	}
	return nil
}
