package base

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go_base_project/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

var instantiation *database = nil
var isDatabaseOpen bool = false

// DbInstantiation Make singleton instantiation
func DbInstantiation() *database {
	if instantiation == nil {
		instantiation = new(database)
	}
	return instantiation
}

type database struct {
	sqlDB  *sql.DB
	gormDB *gorm.DB
}

func (db database) Gorm() *gorm.DB {
	return db.gormDB
}

func (db database) DB() *sql.DB {
	return db.sqlDB
}

func (db database) Close() {
	println("closed database")
	if db.sqlDB != nil {
		db.sqlDB.Close()
		db.sqlDB = nil
		db.gormDB = nil
		isDatabaseOpen = false
	}
}

func OpenDB() *database {
	db := DbInstantiation()

	dbConfig := config.Database()

	if isDatabaseOpen == true {
		return db
	}

	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Name + "?parseTime=" + strings.ToLower(dbConfig.UseTimestamp) + "&loc=" + dbConfig.TimeZone
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
	isDatabaseOpen = true
	return db
}
