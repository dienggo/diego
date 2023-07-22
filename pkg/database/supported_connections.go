package database

import (
	"errors"
	"go_base_project/config"
	"strings"
)

// MYSQL string constanta of supported connection
const MYSQL = "mysql"

// List of supported database connections
var supportedConnections []string = []string{
	MYSQL,
}

// connectionSupport : To check connection supported by string parameter
func connectionSupport(c string) bool {
	var isSupport = false
	for _, value := range supportedConnections {
		if strings.ToLower(value) == c {
			isSupport = true
		}
	}
	return isSupport
}

// checkDatabaseConnectionSupport : To check database configuration connection on supported connections
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
