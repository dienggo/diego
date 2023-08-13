package database

import (
	"github.com/dienggo/diego/config"
	"github.com/dienggo/diego/pkg/logger"
	"gorm.io/gorm"
)

var isConnected = false
var dbConnection *dbc
var mainDsn = ""

type IDbConnection interface {
	Main() *gorm.DB
	Replicas() []*gorm.DB
	Open() *dbc
}

type dbc struct {
	main     *gorm.DB
	replicas []*gorm.DB
}

// Main database connection
func (d dbc) Main() *gorm.DB {
	return d.main
}

// Replicas database connection
func (d dbc) Replicas() []*gorm.DB {
	return d.replicas
}

// Open database connection
func Open() *dbc {
	return openConnection()
}

// Main database uses
func Main() *gorm.DB {
	return Open().Main()
}

// Replicas database uses
func Replicas() []*gorm.DB {
	return Open().Replicas()
}

// Close database connection
func Close() {
	if isConnected == true {
		isConnected = false
		db, err := openConnection().Main().DB()
		if err != nil {
			logger.Error("Close Connection", logger.SetField("error", err.Error()))
			return
		}
		db.Close()
	}
}

// GetMainDsn : getting main dsn
func GetMainDsn() string {
	return mainDsn
}

// openConnection : To open connection with GORM
func openConnection() *dbc {
	err := checkDatabaseConnectionSupport()
	if err != nil {
		logger.Fatal("CheckDatabaseConnectionSupport on openConnection", logger.SetField("error", err.Error()))
		panic(err.Error())
	}

	var db = config.Database()
	if db.Name == "" && db.Username == "" {
		logger.Warn("DB Name & DB Username Empty on openConnection")
		println("[WARNING] App running without database connection")
		return &dbc{}
	}

	var connection *dbc
	switch db.Connection {
	case MYSQL:
		connection = mysqlOpenConnection()
	default:
		connection = mysqlOpenConnection()
	}
	return connection
}
