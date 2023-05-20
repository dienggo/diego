package database

import (
	"database/sql"
	"errors"
	"go_base_project/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

var isConnected = false
var dbConnection *dbc

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

// / To check connection supported by string parameter
func connectionSupport(c string) bool {
	var isSupport = false
	for _, value := range supportedConnections {
		if strings.ToLower(value) == c {
			isSupport = true
		}
	}
	return isSupport
}

// / To check database configuration connection on supported connections
func checkDatabaseConnectionSupport() error {
	var db = config.Database()
	if connectionSupport(db.Connection) == false {
		return errors.New("Unsupported connection for " + db.Connection + ", supported connection are " + strings.Join(supportedConnections, ","))
	}
	for _, replica := range db.Replicas {
		if connectionSupport(replica.Connection) == false {
			return errors.New("Unsupported replica connection for " + replica.Connection + ", supported connection are " + strings.Join(supportedConnections, ","))
		}
	}
	return nil
}

// / Simplify method and reusable method to connect with mysql driver
func connectToMySql(username string, password string, host string, port string, name string, useTimeStamp string) *gorm.DB {
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?parseTime=" + strings.ToLower(useTimeStamp)
	sqlDB, err1 := sql.Open(MYSQL, dsn)

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

	return gormDB
}

// / To open connection with mysql driver
func mysqlOpenConnection() *dbc {
	if isConnected {
		return dbConnection
	}

	var db = config.Database()

	gormDB := connectToMySql(db.Username, db.Password, db.Host, db.Port, db.Name, db.UseTimestamp)

	var replicas []*gorm.DB

	// Do connection on replicas mysql database
	if len(db.Replicas) > 1 {
		for _, replica := range db.Replicas {
			if replica.Connection == MYSQL {
				replicaConnect := connectToMySql(replica.Username, replica.Password, replica.Host, replica.Port, replica.Name, replica.UseTimestamp)
				replicas = append(replicas, replicaConnect)
			}
		}
	}

	// Set initiator variable flag to connected for singleton reason
	isConnected = true
	dbConnection = &dbc{
		main:     gormDB,
		replicas: replicas,
	}

	return dbConnection
}

// / To open connection with GORM
func openConnection() *dbc {
	err := checkDatabaseConnectionSupport()
	if err != nil {
		panic(err.Error())
	}

	var db = config.Database()

	var connection *dbc
	switch db.Connection {
	case MYSQL:
		connection = mysqlOpenConnection()
	default:
		connection = mysqlOpenConnection()
	}
	return connection
}

// Open database connection
func Open() *dbc {
	return openConnection()
}

// Close database connection
func Close() {
	isConnected = false
	db, err := openConnection().Main().DB()
	if err != nil {
		return
	}
	db.Close()

}
