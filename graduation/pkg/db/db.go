package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormDB struct {
	db *gorm.DB
}

func NewGormDB(driver, dsn string) *GormDB {
	var dialector gorm.Dialector
	switch driver {
	case "mysql":
		dialector = mysql.Open(dsn)
	default:
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = db.Debug()

	return &GormDB{db: db}
}
