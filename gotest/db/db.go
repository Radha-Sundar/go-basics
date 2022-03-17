package db

import (
	"gotest/calc"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

var gormConfig = gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
	NowFunc: func() time.Time {
		return calc.Now()
	},
	Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	),
}

func MustInit() {
	var err error
	ConfigDbConnection := "host=localhost user=p_user password=p_pass dbname=p_db port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(ConfigDbConnection), &gormConfig)
	if err != nil {
		panic(err)
	}
}

func Get() *gorm.DB {
	MustInit()
	return db
}
