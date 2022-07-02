package base

import (
	"database/sql"
	"go_base_project/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type database struct {
	sqlDB  *sql.DB
	gormDB *gorm.DB
}

func (db database) Gorm() *gorm.DB {
	return db.gormDB
}

func (db database) Close() {
	db.sqlDB.Close()
}

func OpenDB() database {
	db := database{}

	dbConfig := config.Database()

	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Name + "?parseTime=" + strings.ToLower(dbConfig.UseTimestamp)
	sqlDB, err1 := sql.Open("mysql", dsn)

	if err1 != nil {
		panic(err1.Error())
	}

	gormDB, err2 := gorm.Open(mysql.New(mysql.Config{
		Conn:                     sqlDB,
		DisableDatetimePrecision: false,
	}), &gorm.Config{})

	if err2 != nil {
		panic(err2.Error())
	}

	db.sqlDB = sqlDB
	db.gormDB = gormDB

	return db
}