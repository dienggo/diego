package config

import (
	"github.com/dienggo/diego/pkg/helper"
	"os"
)

var db database
var dbInstance bool = false

type database struct {
	Name                  string
	Username              string
	Password              string
	Host                  string
	Port                  string
	Connection            string
	UseTimestamp          string
	TimeZone              string
	Replicas              []database
	MaxIdleConnection     int
	MaxIdleConnectionTime int // in second
	MaxConnMaxLifetime    int // in second
}

func Database() database {
	if dbInstance == false {
		db = database{
			Name:                  os.Getenv("DB_NAME"),
			Username:              os.Getenv("DB_USERNAME"),
			Password:              os.Getenv("DB_PASSWORD"),
			Host:                  os.Getenv("DB_HOST"),
			Port:                  os.Getenv("DB_PORT"),
			UseTimestamp:          os.Getenv("DB_USE_TIMESTAMP"),
			Connection:            os.Getenv("DB_CONNECTION"),
			TimeZone:              os.Getenv("TIME_ZONE"),
			MaxIdleConnection:     helper.StringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION")),
			MaxIdleConnectionTime: helper.StringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION_TIME")),
			MaxConnMaxLifetime:    helper.StringToInt(os.Getenv("DB_MAX_CONN_MAX_LIFETIME")),
		}
		if helper.IsExistAllEnvKeys("DB_R1_HOST", "DB_R1_PORT", "DB_R1_NAME", "DB_R1_USERNAME", "DB_R1_PASSWORD").Status == true {
			newDB := generateReplica(db, "DB_R1_HOST", "DB_R1_PORT", "DB_R1_NAME", "DB_R1_USERNAME", "DB_R1_PASSWORD")
			db.Replicas = append(db.Replicas, newDB)
		}
		dbInstance = true
	}
	return db
}

func generateReplica(main database, HOST string, PORT string, NAME string, USERNAME string, PASSWORD string) database {
	newDB := main
	newDB.Host = os.Getenv(HOST)
	newDB.Port = os.Getenv(PORT)
	newDB.Name = os.Getenv(NAME)
	newDB.Username = os.Getenv(USERNAME)
	newDB.Password = os.Getenv(PASSWORD)
	return newDB
}
