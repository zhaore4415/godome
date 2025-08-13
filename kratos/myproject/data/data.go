package data

import (
	"time"

	"myproject/internal/proto/hello"

	mgorm "bsi/kratos/micro/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	mgorm.NewGromDB,
)

func NewData(l log.Logger, bs *hello.Bootstrap) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(bs.Data.Database.Source), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	if bs.Service.Env != "prod" {
		db = db.Debug()
	}

	sdb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sdb.SetConnMaxLifetime(30 * time.Minute)
	sdb.SetMaxIdleConns(50)
	sdb.SetMaxOpenConns(100)

	return db, nil
}
